package card

import (
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
	"mesto-goback/internal/service/auth"
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
			card_out := NewCardOutput(h, card)
			output = append(output, card_out)
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

		card_out := NewCardOutput(h, *card2)
		c.JSON(http.StatusCreated, card_out)
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
