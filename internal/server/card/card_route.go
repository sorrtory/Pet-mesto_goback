package card

import (
	"log"
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
	"mesto-goback/internal/server/auth"
	mestoTypes "mesto-goback/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Store *db.Store
}

func NewHTTPHandler(s *db.Store) *HTTPHandler {
	return &HTTPHandler{s}
}

// Return all cards
func (h HTTPHandler) GetCards(c *gin.Context) {
	_, err := auth.Authorized(h.Store, c) // Get user
	if err == nil {
		cards := db.CardGetALL(h.Store)

		output := []CardOutput{}

		// Get likes for every card
		for _, card := range cards {
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
			output = append(output, CardOutput{card, users_liked})
		}

		c.JSON(http.StatusOK, output)
	}
}

// Add new card to database
func (h HTTPHandler) PostCard(c *gin.Context) {
	u, err := auth.Authorized(h.Store, c) // Get user
	if err == nil {
		newCard := mestoTypes.CardPost{}
		if err := c.ShouldBindJSON(&newCard); err != nil {
			// Send 400
			// Can't deserialize
			c.JSON(http.StatusBadRequest, common.FormatValidationError(err))
			return
		}

		// Allow DB to create ID and TIMESTAMP
		card := mestoTypes.Card{Owner_id: u.ID, Name: newCard.Name, Link: newCard.Link}
		card2, err := db.CardPost(h.Store, &card)
		if err != nil {
			errMsg := gin.H{
				"error": gin.H{
					"type":    "database insert",
					"message": err.Error(),
				},
			}

			c.JSON(http.StatusConflict, errMsg)
		}

		if err != nil {
			errMsg := gin.H{
				"error": gin.H{
					"type":    "database select",
					"message": err.Error(),
				},
			}

			c.JSON(http.StatusConflict, errMsg)
			return
		}
		c.JSON(http.StatusCreated, card2)
	}
}

// Delete a card from DB
func (h HTTPHandler) DeleteCard(c *gin.Context) {
	u, err := auth.Authorized(h.Store, c) // Get user
	if err == nil {

		card_id := mestoTypes.CardID{}
		if err := c.ShouldBindUri(&card_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		card, err := db.CardGetByID(h.Store, card_id.CardID)
		if err != nil {
			// Check if card is in db
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if card.Owner_id != u.ID {
			// Prohibit to remove not user's own cards
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "You can delete only your cards"})
			return
		}

		err = db.CardDeleteByID(h.Store, card_id.CardID)
		if err != nil {
			// Can't delete
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"card": "removed"})

	}

}
