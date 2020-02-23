package model

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "db/sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Todo{})
}
