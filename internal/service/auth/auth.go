package auth

import (
	"mesto-goback/internal/db"
	mestoTypes "mesto-goback/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check for Authorization header == User password
//
// Response an error if wrong password
func Authorized(store *db.Store, c *gin.Context) (*mestoTypes.User, error) {
	auth := mestoTypes.UserAuth{}
	if err := c.ShouldBindHeader(&auth); err != nil {
		// Hasn't Authorization header
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	u, err := db.UserGetByPassword(store, auth)
	if err != nil {
		// Authorization isn't found in DB
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return nil, err
	}
	return u, nil
}
