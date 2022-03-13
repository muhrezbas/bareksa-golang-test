package mysql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// LOC localtime godoc
const LOC = "Local"

// Config godoc
type Config struct {
	Host string
	DB   string
	User string
	Pass string
	Port int
}

// Connect godoc
func (conf Config) Connect() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.DB, LOC,
	)

	db, err = gorm.Open("mysql", DBURL)
	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		db.Close()
	} else {
		log.Println("MySQL Has Connected ............................!")
		tables := []string{}
		db.Select(&tables, "SHOW TABLES")
	}
	db.DB().SetMaxIdleConns(0)
	return db, nil
}
