package api

import (
	"mesto-goback/internal/db"
	"mesto-goback/internal/server/card"
	"mesto-goback/internal/server/user"

	"github.com/gin-gonic/gin"
)

type API struct {
	address string
	port    string
	store   *db.Store
}

func NewAPI(address string, port string, store *db.Store) *API {
	return &API{address, port, store}
}

// Start handlers
func (a API) Serve() {
	router := gin.Default()

	usersMeRoute := router.Group("/users/me")
	{
		userHandler := user.NewHTTPHandler(a.store)
		usersMeRoute.GET("/", userHandler.GetMe)
		usersMeRoute.PATCH("/", userHandler.PatchMe)
		usersMeRoute.PATCH("/avatar", userHandler.PatchMeAvatar)

	}
	cardsRoute := router.Group("/cards")
	{
		cardsHandler := card.NewHTTPHandler(a.store)
		cardsRoute.GET("/", cardsHandler.GetCards)
		cardsRoute.POST("/", cardsHandler.PostCards)
	}

	router.Run(a.address + ":" + a.port)
}
