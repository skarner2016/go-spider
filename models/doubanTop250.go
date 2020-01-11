package models

import (
	"time"
)

type DoubanTop250 struct {
	ID        uint   `gorm:"primary_key auto_increment"`
	Sort      int    `gorm:"index:sort"`
	Name      string `gorm:"index:name"`
	Star      string
	Number    int
	Desc      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (d *DoubanTop250) Migrate() {
	db.AutoMigrate(&DoubanTop250{})
}

func CreateDoubanTop250(d *DoubanTop250) error {
	return db.Create(d).Error
}
