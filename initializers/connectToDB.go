package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func ConnectToDB() {
	if DB == nil {
		dsn := os.Getenv("DB_DNS")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		DB = db
	}
}
