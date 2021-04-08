package models
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&Product{})
	DB=database
}
