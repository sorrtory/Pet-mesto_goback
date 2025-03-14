package web

import (
	"mesto-goback/internal/common"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h HTTPHandler) ServeFrontEnd(route string, router *gin.Engine) {
	PUBLIC := common.GetEnv("BACKEND_PUBLIC")
	router.Static(route, PUBLIC)
}

// To load a single file

// func (h HTTPHandler) LoadHTML(router *gin.Engine) {
// 	router.LoadHTMLGlob("internal/frontend/public/*")
// }

// func (h HTTPHandler) GetIndex(c *gin.Context) {
// 	c.HTML(http.StatusOK, "index.html", gin.H{
// 		"title": "Main website",
// 	})
// }
