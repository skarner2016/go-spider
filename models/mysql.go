package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_TYPE = "mysql"
	DB_USER = "root"
	DB_PASS = "123456"
	DB_HOST = "127.0.0.1:3306"
	DB_NAME = "go_spider"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(DB_TYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_NAME))

	if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
