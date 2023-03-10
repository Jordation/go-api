package orm

import (
	"log"
)

func CheckMatchLink(url string) error {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}

	if err := db.Create(&HitLink{Link: url}).Error; err != nil {
		log.Println("Error adding new url to db: ", err)
		return err
	}
	return nil
}
