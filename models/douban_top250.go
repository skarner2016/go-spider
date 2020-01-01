package models

import (
	"time"
)

type DoubanTop250 struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"index:name"`
	Star      string
	Number    string
	Desc      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func Test() {
	//db.AutoMigrate(&DoubanTop250{})
	//fmt.Println("Test", db)
}

func CreateDoubanTop250(d *DoubanTop250) error {
	return db.Create(d).Error
}
