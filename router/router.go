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

	// bareksaHandler := handler.Context{}

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
	// 	}
	// }
	return
}
