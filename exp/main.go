package main

import (
	"Gogram/models"
	"fmt"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "gogram_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	user := models.User{
		Name:  "lexville",
		Email: "lexville.io",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	user.Email = "lexville1.io"
	if err := us.Update(&user); err != nil {
		panic(err)
	}
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	userByEmail, err := us.ByEmail("lexville1.io")
	if err != nil {
		panic(err)
	}
	fmt.Println(userByEmail)
}
