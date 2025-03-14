package like

import (
	"mesto-goback/internal/db"
	"mesto-goback/internal/service/auth"
	mestoTypes "mesto-goback/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	store *db.Store
}

func NewHTTPHandler(store *db.Store) *HTTPHandler {
	return &HTTPHandler{store}
}

func (h HTTPHandler) PutLike(c *gin.Context) {
	u, err := auth.Authorized(h.store, c)
	if err == nil {
		card_id := mestoTypes.CardID{}
		if err := c.ShouldBindUri(&card_id); err != nil {
			// Bad cardId format
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

        // Check like existance
        _, err := db.LikeExists(h.store, u.ID, card_id.CardID)
        if err == nil {
			c.JSON(http.StatusConflict, gin.H{"like": "exists"})
            return
        }

        // Set like
        _, err = db.LikeSetByCardID(h.store, u.ID, card_id.CardID)
        if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"like": "set"})
	}
}

func (h HTTPHandler) DeleteLike(c *gin.Context) {
	u, err := auth.Authorized(h.store, c)
	if err == nil {
		card_id := mestoTypes.CardID{}
		if err := c.ShouldBindUri(&card_id); err != nil {
			// Bad cardId format
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

        // Check like existance
        _, err := db.LikeExists(h.store, u.ID, card_id.CardID)
        if err != nil {
			c.JSON(http.StatusConflict, gin.H{"like": "doesn't exist"})
            return
        }

        // Remove like
        _, err = db.LikeDeleteByCardID(h.store, u.ID, card_id.CardID)
        if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"like": "removed"})
	}
}
