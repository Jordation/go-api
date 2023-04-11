package db

import (
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Adapter struct {
	DB *gorm.DB
}

func GetDbPath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, "new_DEV.db")
	return filePath
}
func Getdb() *Adapter {
	db, err := gorm.Open(sqlite.Open(GetDbPath()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Adapter{
		db,
	}
}
