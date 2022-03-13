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

// Rtx godoc
type Rtx interface {
	CreateNews(body News) (News, error)
	FindNews() ([]News, error)
	FindNewsByID(id string) (News, error)
	FindNewsByTopic(id string) (News, error)
	FindNewsByStatus(id string) (News, error)
	UpdateNews(body News, id string) (News, error)
}

// Pengguna godoc
type News struct {
	ID        int64      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not_null"`
	Name      string     `json:"name" gorm:"column:name"`
	TopicID   int        `json:"topic_id" gorm:"column:topic_id"`
	Status    string     `json:"status" gorm:"column:status"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewWathdog godoc
func NewNews(db *gorm.DB, redisCache *redis.Client) Rtx {
	return &news{db: db, redisCache: redisCache}
}

type news struct {
	db         *gorm.DB
	redisCache *redis.Client
}

func (watch *news) CreateNews(body News) (News, error) {
	fmt.Println(body.Name, "haloooo news")
	result := watch.db.Table("news").Create(&body)
	fmt.Println(result)
	return body, nil
}
func (watch *news) UpdateNews(body News, id string) (News, error) {
	fmt.Println(body.Name, "haloooo")
	// result := watch.db.Table("news").Create(&body)
	result := watch.db.Table("news").Where("id = ?", id).Updates(map[string]interface{}{"name": body.Name})
	fmt.Println(result)
	return body, nil
}
func (watch *news) FindNews() ([]News, error) {

	val, err := watch.redisCache.Get("newsAll").Result()
	if err != nil {
		var dirs []News
		if err := watch.db.Table("news").Find(&dirs).Error; err != nil {
			return dirs, err
		}
		p, err := json.Marshal(dirs)
		if err != nil {
			return dirs, err
		}
		tess := watch.redisCache.Set("newsAll", p, 8000000000).Err()
		if tess != nil {
			return dirs, err
		}
		return dirs, nil
	}
	deserialized := []News{}
	tess := json.Unmarshal([]byte(val), &deserialized)
	if tess == nil {
		return deserialized, nil
	}
	return deserialized, tess
}

func (watch *news) FindNewsByStatus(id string) (News, error) {
	myKey := "newsStatus" + id
	fmt.Println(myKey, "tess")
	val, err := watch.redisCache.Get(myKey).Result()
	if err != nil {
		fmt.Println(err, "error")
		var dirs News
		if err := watch.db.Table("news").Where("`status` = ?", id).Find(&dirs).Error; err != nil {
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
	var deserialized News
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
func (watch *news) FindNewsByID(id string) (News, error) {
	myKey := "news" + id
	fmt.Println(myKey, "tess")
	val, err := watch.redisCache.Get(myKey).Result()
	if err != nil {
		fmt.Println(err, "error")
		var dirs News
		if err := watch.db.Table("news").Where("`id` = ?", id).Find(&dirs).Error; err != nil {
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
	var deserialized News
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
func (watch *news) FindNewsByTopic(id string) (News, error) {
	myKey := "newsTopic" + id
	fmt.Println(myKey, "tess")
	val, err := watch.redisCache.Get(myKey).Result()
	if err != nil {
		fmt.Println(err, "error")
		var dirs News
		if err := watch.db.Table("news").Where("`topic_id` = ?", id).Find(&dirs).Error; err != nil {
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
	var deserialized News
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
