package db

import (
	"database/sql"
	"log"
	mestoTypes "mesto-goback/internal/types"
)

// Get array of likes by card ID
func LikesGetByCardID(store *Store, card_id int) []mestoTypes.Like {
	query := `SELECT * FROM likes WHERE card_id=$1`
	var likes []mestoTypes.Like

	rows, err := store.DB.Query(query, card_id)
	if err != nil {
		return likes
	}
	defer rows.Close()

	for rows.Next() {
		card, err := LikeScan(rows)
		if err != nil {
			log.Printf("No cards %v\n", err)
			return likes
		}
		likes = append(likes, *card)
	}
	return likes
}

// Select like from db
func LikeExists(store *Store, user_id int, card_id int) (*mestoTypes.Like, error) {
    query := `SELECT * FROM likes WHERE user_id=$1 AND card_id=$2`
    r := store.DB.QueryRow(query, user_id, card_id)
    return LikeScan(r)
}

// Set like to a card
func LikeSetByCardID(store *Store, user_id int, card_id int) (*mestoTypes.Like, error){
    query := `INSERT INTO likes (user_id, card_id) VALUES ($1, $2) RETURNING * `
    // query := `INSERT INTO likes (user_id, card_id) VALUES ($1, $2) RETURNING id, user_id, card_id`
    r := store.DB.QueryRow(query, user_id, card_id)
    return LikeScan(r)
}

// Delete like from a card
func LikeDeleteByCardID(store *Store, user_id int, card_id int) (*mestoTypes.Like, error) {
    query := `DELETE FROM likes WHERE user_id=$1 AND card_id=$2 RETURNING *`
    r := store.DB.QueryRow(query, user_id, card_id)
    return LikeScan(r)
}


// Scan one like from a row
func LikeScan(row_s interface{ Scan(args ...any) error }) (*mestoTypes.Like, error) {
	like := mestoTypes.Like{}
	err := row_s.Scan(&like.ID, &like.User_ID, &like.Card_ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil, err
	}
	return &like, nil
}
