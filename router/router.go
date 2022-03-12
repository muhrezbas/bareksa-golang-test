package router

import (
	// "os"

	"github.com/gin-gonic/gin"
)

// Context godoc
type Context struct {
	R *gin.Engine
}

// LoadRoutes godoc
func (route *Context) LoadRoutes() {

	// gfxHandler := handler.Context{}

	route.R.GET("/ping", func(c *gin.Context) {
		ping := map[string]interface{}{
			"status":   true,
			"response": "pong",
		}
		c.JSONP(200, ping)
		// return
	})
	// API ENDPOINT
	// api := route.R.Group("/api")
	// {
	// 	{
	// 		api.GET("/", gfxHandler.Hello)
	// 		api.GET("/palindrome", gfxHandler.Palindrome)
	// 		api.GET("/language", gfxHandler.Language)
	// 		api.POST("/language", gfxHandler.PostLanguage)
	// 		api.GET("/languages", gfxHandler.GetAllLanguages)
	// 		api.GET("/language/:id", gfxHandler.GetOneLanguage)
	// 		api.PATCH("/language/:id", gfxHandler.UpdateOneLanguage)
	// 		api.DELETE("/language/:id", gfxHandler.DeleteOneLanguage)
	// 	}
	// }
	return
}
