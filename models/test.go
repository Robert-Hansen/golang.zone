package models

import (
	"time"
)

type PostTest struct {
	ID       int    `gorm:"primary_key"`
	Title    string `gorm:"size:255;unique"`
	slug     string `gorm:"size:255;unique"`
	AuthorID int
	Author   *User `gorm:"ForeignKey:AuthorID"`
	Category *Category
	//PublishedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"`
}

func (PostTest) TableName() string {
	return "posts"
}
