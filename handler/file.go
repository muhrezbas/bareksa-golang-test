package handler

import (
	"bareksa-api/model"
	"bareksa-api/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Context godoc
type Context struct {
	UC usecase.Context
}
type GetTags struct {
	NewsID []int  `json:"news_id"`
	Name   string `json:"name"`
}

// GetStatus godoc
func (h *Context) GetTopic(c *gin.Context) {

	result, err := h.UC.GetTopic()
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetTopicByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := h.UC.GetTopicByID(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}

func (h *Context) CreateTopic(c *gin.Context) {
	var topicData model.Topic
	if err := c.ShouldBindJSON(&topicData); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(topicData, "handler topic")
	result, err := h.UC.CreateTopic(topicData)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}

func (h *Context) UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	var topicData model.Topic
	if err := c.ShouldBindJSON(&topicData); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(topicData, "handler")
	result, err := h.UC.UpdateTopic(topicData, id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) CreateNews(c *gin.Context) {
	var newsData model.News
	if err := c.ShouldBindJSON(&newsData); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(newsData, "handler news")
	result, err := h.UC.CreateNews(newsData)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetNews(c *gin.Context) {

	result, err := h.UC.GetNews()
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := h.UC.GetNewsByID(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetNewsByTopic(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := h.UC.GetNewsByTopic(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetNewsByStatus(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := h.UC.GetNewsByStatus(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}

func (h *Context) UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var newsData model.News
	if err := c.ShouldBindJSON(&newsData); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(newsData, "handler")
	result, err := h.UC.UpdateNews(newsData, id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) CreateTags(c *gin.Context) {
	// var tagsData model.Tags
	// var tagsNewsData model.TagsNews
	var getTags GetTags
	if err := c.ShouldBindJSON(&getTags); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(getTags, "fafmoa")
	tags := model.Tags{Name: getTags.Name}
	resultTags, err := h.UC.CreateTags(tags)
	fmt.Println(resultTags, "handler tagss", err)
	for _, f := range getTags.NewsID {
		fmt.Println(f)
		user := model.TagsNews{NewsID: f, TagID: resultTags.ID}
		result, err := h.UC.CreateTagsNews(user)
		fmt.Println(result, "handler tagss news", err)
	}

	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, resultTags)
	return
}
func (h *Context) GetTags(c *gin.Context) {

	result, err := h.UC.GetTags()
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
func (h *Context) GetTagsByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := h.UC.GetTagsByID(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}

func (h *Context) UpdateTags(c *gin.Context) {
	id := c.Param("id")
	var newsData model.Tags
	if err := c.ShouldBindJSON(&newsData); err != nil {
		SendBadRequets(c, err.Error())
		return
	}
	fmt.Println(newsData, "handler")
	result, err := h.UC.UpdateTags(newsData, id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	SendSuccess(c, result)
	return
}
