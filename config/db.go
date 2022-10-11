package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	e "hacktiv-assignment-2/entity"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pass = "1234"
	dbName = "e-commerce"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", 
		host, port, user, pass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { panic(err) }
	fmt.Printf("%s\n\n", "Successfully connected to database!")
	db.AutoMigrate(&e.Orders{}, &e.Items{})
	DB = db
}
