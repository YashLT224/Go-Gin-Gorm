package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	const DNS = "root:<>@tcp(127.0.0.1:3306)/media?parseTime=true"

	return gorm.Open(mysql.Open(DNS), &gorm.Config{})
}
