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

// // Get one card by SQL query
// func CardGet(store *Store, query string, args ...any) (*mestoTypes.Card, error) {
// 	row := store.DB.QueryRow(query, args...)
// 	return CardScan(row)
// }

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
