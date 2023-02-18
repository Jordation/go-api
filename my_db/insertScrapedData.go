package my_db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InsertData(e Event) {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Where(Event{EventName: e.EventName}).FirstOrCreate(&e)
	if result.RowsAffected == 0 {
		var existingEvent Event
		db.Where("event_name", e.EventName).Find(&existingEvent)
		eID := existingEvent.ID
		for _, v := range e.Maps {
			v.EventID = eID
			db.Create(&v)
		}
	}
}
