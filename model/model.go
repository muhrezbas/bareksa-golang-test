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

// Ctx godoc
type Ctx interface {
	FindTopic() ([]Topic, error)
	FindTopicByID(id string) (Topic, error)
	CreateTopic(body Topic) (Topic, error)
	UpdateTopic(body Topic, id string) (Topic, error)
}

// Pengguna godoc
type Topic struct {
	ID        int64      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not_nuull"`
	Name      string     `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewWathdog godoc
func NewWathdog(db *gorm.DB, redisCache *redis.Client) Ctx {
	return &watching{db: db, redisCache: redisCache}
}

type watching struct {
	db         *gorm.DB
	redisCache *redis.Client
}

func (watch *watching) UpdateTopic(body Topic, id string) (Topic, error) {
	fmt.Println(body.Name, "haloooo")
	// result := watch.db.Table("topic").Create(&body)
	result := watch.db.Table("topic").Where("id = ?", id).Updates(map[string]interface{}{"name": body.Name})
	fmt.Println(result)
	return body, nil
}
func (watch *watching) CreateTopic(body Topic) (Topic, error) {
	fmt.Println(body.Name, "haloooo")
	result := watch.db.Table("topic").Create(&body)
	fmt.Println(result)
	return body, nil
}

func (watch *watching) FindTopic() ([]Topic, error) {
	fmt.Println("model")

	val, err := watch.redisCache.Get("topicAll").Result()
	fmt.Println(val, err, "oajpv")
	if err != nil {
		fmt.Println("model non redis")
		var dirs []Topic
		if err := watch.db.Table("topic").Find(&dirs).Error; err != nil {
			return dirs, err
		}
		p, err := json.Marshal(dirs)
		if err != nil {
			return dirs, err
		}
		tess := watch.redisCache.Set("topicAll", p, 8000000000).Err()
		if tess != nil {
			return dirs, err
		}
		return dirs, nil
	}
	fmt.Println("with redis")
	deserialized := []Topic{}
	tess := json.Unmarshal([]byte(val), &deserialized)
	if tess == nil {
		return deserialized, nil
	}
	return deserialized, tess
}

func (watch *watching) FindTopicByID(id string) (Topic, error) {
	myKey := "topic" + id
	fmt.Println(myKey, "tess")
	val, err := watch.redisCache.Get(myKey).Result()
	if err != nil {
		fmt.Println(err, "error")
		var dirs Topic
		if err := watch.db.Table("topic").Where("`id` = ?", id).Find(&dirs).Error; err != nil {
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
	var deserialized Topic
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
