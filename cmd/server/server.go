package server

import (
	"mesto-goback/internal/db"
	"mesto-goback/web"
	"mesto-goback/internal/service/card"
	"mesto-goback/internal/service/like"
	"mesto-goback/internal/service/user"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	port    string
	store   *db.Store
}

func NewServer(address string, port string, store *db.Store) *Server {
	return &Server{address, port, store}
}

// Set CORS headers for all responses
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

// Start handlers
func (a Server) Serve() {
	router := gin.Default()
    router.Use(CORSMiddleware())

	// Serve frontend
	front := web.NewHTTPHandler()
	front.ServeFrontEnd("/static", router)

	// Serve api
	apiRoute := router.Group("/api")
	{
		// Serve users api
		usersMeRoute := apiRoute.Group("/users/me")
		{
			userHandler := user.NewHTTPHandler(a.store)
			usersMeRoute.GET("/", userHandler.GetMe)
			usersMeRoute.PATCH("/", userHandler.PatchMe)
			usersMeRoute.PATCH("/avatar", userHandler.PatchMeAvatar)
		}

		// Serve cards api
		cardsRoute := apiRoute.Group("/cards")
		{
			cardsHandler := card.NewHTTPHandler(a.store)
			cardsRoute.GET("/", cardsHandler.GetCards)
			cardsRoute.POST("/", cardsHandler.PostCard)
			cardsRoute.DELETE("/:card_id", cardsHandler.DeleteCard)

			likesRoute := cardsRoute.Group("/likes")
			{
				likeHandler := like.NewHTTPHandler(a.store)
				likesRoute.PUT("/:card_id", likeHandler.PutLike)
				likesRoute.DELETE("/:card_id", likeHandler.DeleteLike)
			}
		}
	}

	router.Run(a.address + ":" + a.port)
}
