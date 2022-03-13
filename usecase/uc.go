package usecase

import (
	"bareksa-api/model"

	// "time.Time "

	// "bareksa-api/pkg/config"
	"bareksa-api/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Context godoc
type Context struct {
	DB         *gorm.DB
	RedisCache *redis.Client
	Config     config.Interface
}
type Request struct {
	access_token string `json:"access_token"`
	username     string `json:"username"`
}

func (uc *Context) CreateTagsNews(body model.TagsNews) (model.TagsNews, error) {
	mdl := model.NewTags(uc.DB, uc.RedisCache)
	return mdl.CreateTagsNews(body)
}

func (uc *Context) GetTopic() ([]model.Topic, error) {
	mdl := model.NewWathdog(uc.DB, uc.RedisCache)
	// fmt.Println("fejpiwfasjfccjcj")
	return mdl.FindTopic()
}
func (uc *Context) GetTopicByID(id string) (model.Topic, error) {
	mdl := model.NewWathdog(uc.DB, uc.RedisCache)
	return mdl.FindTopicByID(id)
}
func (uc *Context) CreateTopic(body model.Topic) (model.Topic, error) {
	mdl := model.NewWathdog(uc.DB, uc.RedisCache)
	return mdl.CreateTopic(body)
}
func (uc *Context) UpdateTopic(body model.Topic, id string) (model.Topic, error) {
	mdl := model.NewWathdog(uc.DB, uc.RedisCache)
	return mdl.UpdateTopic(body, id)
}
func (uc *Context) CreateNews(body model.News) (model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.CreateNews(body)
}
func (uc *Context) GetNews() ([]model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.FindNews()
}
func (uc *Context) GetNewsByID(id string) (model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.FindNewsByID(id)
}
func (uc *Context) GetNewsByTopic(id string) (model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.FindNewsByTopic(id)
}
func (uc *Context) GetNewsByStatus(id string) (model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.FindNewsByStatus(id)
}
func (uc *Context) UpdateNews(body model.News, id string) (model.News, error) {
	mdl := model.NewNews(uc.DB, uc.RedisCache)
	return mdl.UpdateNews(body, id)
}
func (uc *Context) CreateTags(body model.Tags) (model.Tags, error) {
	mdl := model.NewTags(uc.DB, uc.RedisCache)
	return mdl.CreateTags(body)
}
func (uc *Context) GetTags() ([]model.Tags, error) {
	mdl := model.NewTags(uc.DB, uc.RedisCache)
	return mdl.FindTags()
}
func (uc *Context) GetTagsByID(id string) (model.Tags, error) {
	mdl := model.NewTags(uc.DB, uc.RedisCache)
	return mdl.FindTagsByID(id)
}
func (uc *Context) UpdateTags(body model.Tags, id string) (model.Tags, error) {
	mdl := model.NewTags(uc.DB, uc.RedisCache)
	return mdl.UpdateTags(body, id)
}
