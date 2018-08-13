package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "gogram_dev"
)

// User contains a name and email
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.DropTableIfExists(&User{})
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "lexville", Email: "lexville.io"})

	var user User
	db.First(&user, 1)
	fmt.Println(user) // find product with id 1
	db.First(&user, "email = ?", "lexville.io")
	fmt.Println(user)
}
