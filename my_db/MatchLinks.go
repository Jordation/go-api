package my_db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CheckMatchLink(url string) error {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.Create(&HitLink{Link: url}).Error; err != nil {
		log.Println("Error adding new url to db: ", err)
		return err
	}
	return nil
}
