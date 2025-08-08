package db

import (
	"fmt"
	"log"
	"scheduler/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", models.Config.Mysql.User, models.Config.Mysql.Password, models.Config.Mysql.Host, models.Config.Mysql.Port, models.Config.Mysql.Name)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	for table, query := range models.Config.Tables {
		err = DB.Exec(query).Error
		if err != nil {
			log.Fatal("Failed to create table:", table, err)
		} else {
			log.Println("Table created or exists:", table)
		}
	}
}
