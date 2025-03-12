package db

import (
	"database/sql"
	"log"
	mestoTypes "mesto-goback/internal/types"
)

func CardGetALL(store *Store) []mestoTypes.Card {
	query := `SELECT * FROM cards`
	var cards []mestoTypes.Card

	rows, err := store.DB.Query(query)
	if err != nil {
		return cards
	}
	defer rows.Close()

	for rows.Next() {
		card, err := CardScan(rows)
		if err != nil {
			log.Printf("No cards %v\n", err)
			return cards
		}
		cards = append(cards, *card)
	}
	return cards
}

// Delete card by its ID
func CardDeleteByID(store *Store, card_id int) error {
	query := `DELETE FROM cards WHERE id=$1`
    _, err := store.DB.Exec(query, card_id)
    return err
}

// Insert card data into database. Return assigned ID
func CardPost(store *Store, card *mestoTypes.Card) (*mestoTypes.Card, error) {
	query := `INSERT INTO cards (owner_id, name, link) VALUES ($1, $2, $3) RETURNING id, owner_id, name, link, createdAt`
	return CardGet(store, query, card.Owner_id, card.Name, card.Link)
}

// Get one card from DB by its ID
func CardGetByID(store *Store, id int) (*mestoTypes.Card, error) {
	query := `SELECT * FROM cards WHERE id=$1`
	return CardGet(store, query, id)
}

// Get one card by SQL query
func CardGet(store *Store, query string, args ...any) (*mestoTypes.Card, error) {
	row := store.DB.QueryRow(query, args...)
	return CardScan(row)
}

// Scan card from a row
func CardScan(row_s interface{ Scan(args ...any) error }) (*mestoTypes.Card, error) {
	card := mestoTypes.Card{}
	err := row_s.Scan(&card.ID, &card.Owner_id, &card.Name, &card.Link, &card.CreatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil, err
	}
	return &card, nil
}
