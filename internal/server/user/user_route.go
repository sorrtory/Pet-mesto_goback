package user

import (
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
	"mesto-goback/internal/server/auth"
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

// Send User to user
func (h HTTPHandler) GetMe(c *gin.Context) {
	u, err := auth.Authorized(h.store, c) // Get user
	if err == nil {
		c.JSON(http.StatusOK, *u)
	}
}

func (h HTTPHandler) PatchMeAvatar(c *gin.Context) {
	u, err := auth.Authorized(h.store, c) // Get user
	if err == nil {
		avatar := mestoTypes.UserAvatar{}
		if err := c.ShouldBindJSON(&avatar); err != nil {
			// Can't deserialize
			c.JSON(http.StatusBadRequest, common.FormatValidationError(err))
			return
		}

		err := db.UserUpdateAvatar(h.store, *u, avatar)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		// Get user from DB again
		u, err = db.UserGetByID(h.store, u.ID)
		if err != nil {
			// Can't get user from DB
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, *u)

	}

}

// Update User's name and about
func (h HTTPHandler) PatchMe(c *gin.Context) {
	u, err := auth.Authorized(h.store, c) // Get user
	if err == nil {
		me := mestoTypes.UserMe{}
		if err := c.ShouldBindJSON(&me); err != nil {
			// Can't deserialize
			c.JSON(http.StatusBadRequest, common.FormatValidationError(err))
			return
		}

		// Update user
		db.UserRenameMe(h.store, *u, me)

		// Get user from DB again
		u, err = db.UserGetByID(h.store, u.ID)
		if err != nil {
			// Can't get user from DB
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, *u)
	}
}
