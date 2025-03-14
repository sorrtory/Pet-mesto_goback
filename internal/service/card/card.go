package card

import (
	"log"
	"mesto-goback/internal/db"
	mestoTypes "mesto-goback/internal/types"
)

type CardOutput struct {
	mestoTypes.Card
	Likes []mestoTypes.User `json:"likes"`
}

// Add likes to to card. Replace user_id with users
func NewCardOutput(h HTTPHandler, card mestoTypes.Card) CardOutput{

    var users_liked []mestoTypes.User
    likes := db.LikesGetByCardID(h.Store, card.ID)
    // Get user for every like
    for _, like := range likes {
        user, err := db.UserGetByID(h.Store, like.User_ID)
        if err != nil {
            log.Printf("Card's user not found: %v\n", err.Error())
        }
        users_liked = append(users_liked, *user)
    }

    return CardOutput{card, users_liked}
}
