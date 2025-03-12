package mestoTypes

import "time"

type Card struct {
	ID        int       `json:"id"`
	Owner_id  int       `json:"owner_id"`
	Name      string    `json:"string"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

type CardPost struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}
