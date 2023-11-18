package models

import (
	"github.com/furkanadiiguzel/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
