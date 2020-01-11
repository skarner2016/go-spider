package models

import "time"

// TODO:
type LagouChengdu struct {
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
