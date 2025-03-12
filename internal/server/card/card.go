package card

import "mesto-goback/internal/types"

type CardOutput struct {
	mestoTypes.Card
	Likes []mestoTypes.User `json:"likes"`
}
