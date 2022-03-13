package router

import (
	// "os"
	"bareksa-api/config"
	"bareksa-api/handler"
	"bareksa-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Context godoc
type Context struct {
	R          *gin.Engine
	DB         *gorm.DB
	Config     config.Interface
	RedisCache *redis.Client
}

// LoadRoutes godoc
func (route *Context) LoadRoutes() {

	bareksaHandler := handler.Context{
		UC: usecase.Context{
			DB:         route.DB,
			Config:     route.Config,
			RedisCache: route.RedisCache,
		},
	}

	route.R.GET("/ping", func(c *gin.Context) {
		ping := map[string]interface{}{
			"status":   true,
			"response": "pong",
		}
		c.JSONP(200, ping)
		// return
	})
	// APsetupRouterI ENDPOINT
	{
		api := route.R.Group("/api")
		{
			api.POST("/news", bareksaHandler.CreateNews)
			api.GET("/news", bareksaHandler.GetNews)
			api.GET("/news/:id", bareksaHandler.GetNewsByID)
			api.PUT("/news/:id", bareksaHandler.UpdateNews)
			api.GET("/news-topic/:id", bareksaHandler.GetNewsByTopic)
			api.GET("/news-status/:id", bareksaHandler.GetNewsByStatus)
			api.GET("/topic", bareksaHandler.GetTopic)
			api.GET("/topic/:id", bareksaHandler.GetTopicByID)
			api.POST("/topic", bareksaHandler.CreateTopic)
			api.PUT("/topic/:id", bareksaHandler.UpdateTopic)
			api.POST("/tags", bareksaHandler.CreateTags)
			api.GET("/tags", bareksaHandler.GetTags)
			api.GET("/tags/:id", bareksaHandler.GetTagsByID)
			api.PUT("/tags/:id", bareksaHandler.UpdateTags)
		}
	}
	return
}
