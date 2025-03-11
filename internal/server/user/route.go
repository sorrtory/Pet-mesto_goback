package user

import (
	"mesto-goback/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store *db.Store
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{store}
}

func (h Handler) Authorized(c *gin.Context) (*User, error) {
	auth := UserAuth{}
	if err := c.ShouldBindHeader(&auth); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return nil, err
	}
	u, err := GetUser(h.store, auth)
	return u, err

}

func (h Handler) Me(c *gin.Context) {
	u, err := h.Authorized(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, *u)
	}

}
