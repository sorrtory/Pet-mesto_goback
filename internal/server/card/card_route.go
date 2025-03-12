package card

import (
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
		c.JSON(http.StatusOK, cards)
	}
}

func (h HTTPHandler) PostCards(c *gin.Context) {
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
