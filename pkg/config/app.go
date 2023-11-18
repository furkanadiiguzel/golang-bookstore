package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "furkan:...")
	if err != nil {
		panic(err)
	}
	db = d
}
func getDB() *gorm.DB {
	return db
}
