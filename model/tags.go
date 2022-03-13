package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	// "strconv"
	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// Ttx godoc
type Ttx interface {
	CreateTags(body Tags) (Tags, error)
	FindTags() ([]Tags, error)
	FindTagsByID(id string) (Tags, error)
	UpdateTags(body Tags, id string) (Tags, error)
	CreateTagsNews(body TagsNews) (TagsNews, error)
}

// Pengguna godoc
type Tags struct {
	ID        int64      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not_null"`
	Name      string     `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type TagsNews struct {
	ID        int64      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not_null"`
	NewsID    int        `json:"news_id" gorm:"column:news_id"`
	TagID     int64      `json:"tag_id" gorm:"column:tag_id"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

// NewWathdog godoc
func NewTags(db *gorm.DB, redisCache *redis.Client) Ttx {
	return &tags{db: db, redisCache: redisCache}
}

type tags struct {
	db         *gorm.DB
	redisCache *redis.Client
}

func (watch *tags) CreateTagsNews(body TagsNews) (TagsNews, error) {
	fmt.Println(body, "feafno")
	result := watch.db.Table("tags_news").Create(&body)
	fmt.Println(result)
	return body, nil
}

func (watch *tags) CreateTags(body Tags) (Tags, error) {

	result := watch.db.Table("tags").Create(&body)
	fmt.Println(result)
	return body, nil
}
func (watch *tags) UpdateTags(body Tags, id string) (Tags, error) {
	fmt.Println(body.Name, "haloooo")
	// result := watch.db.Table("tags").Create(&body)
	result := watch.db.Table("tags").Where("id = ?", id).Updates(map[string]interface{}{"name": body.Name})
	fmt.Println(result)
	return body, nil
}
func (watch *tags) FindTags() ([]Tags, error) {

	val, err := watch.redisCache.Get("tagsAll").Result()
	if err != nil {
		var dirs []Tags
		if err := watch.db.Table("tags").Find(&dirs).Error; err != nil {
			return dirs, err
		}
		p, err := json.Marshal(dirs)
		if err != nil {
			return dirs, err
		}
		tess := watch.redisCache.Set("tagsAll", p, 8000000000).Err()
		if tess != nil {
			return dirs, err
		}
		return dirs, nil
	}
	deserialized := []Tags{}
	tess := json.Unmarshal([]byte(val), &deserialized)
	if tess == nil {
		return deserialized, nil
	}
	return deserialized, tess
}

func (watch *tags) FindTagsByID(id string) (Tags, error) {
	myKey := "tags" + id
	fmt.Println(myKey, "tess")
	val, err := watch.redisCache.Get(myKey).Result()
	if err != nil {
		fmt.Println(err, "error")
		var dirs Tags
		if err := watch.db.Table("tags").Where("`id` = ?", id).Find(&dirs).Error; err != nil {
			return dirs, err
		}
		p, err := json.Marshal(dirs)
		if err != nil {
			return dirs, err
		}
		tess := watch.redisCache.Set(myKey, p, 8000000000).Err()
		if tess != nil {
			fmt.Println(tess, "err redis")
			return dirs, err
		}
		return dirs, nil
	}
	fmt.Println("with redis")
	var deserialized Tags
	rawIn := json.RawMessage(val)

	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	tess := json.Unmarshal(bytes, &deserialized)
	if tess == nil {
		return deserialized, nil
	}
	return deserialized, tess
}
