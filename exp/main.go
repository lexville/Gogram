package main

import (
	"fmt"
	"html/template"
	"os"
)

type Job struct {
	Occupation string
}

type User struct {
	Name     string
	Nickname string
	Dog      map[string]string
	Job      Job
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := User{
		Name:     "Lexville",
		Nickname: "buddy",
		Dog: map[string]string{
			"Name": "bobo",
			"Age":  "2",
		},
		Job: Job{Occupation: "Dentist"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Sally"
	data.Nickname = "Kitty"
	data.Dog = map[string]string{"Name": "zulo", "Age": "1"}
	data.Job.Occupation = "Nurse"

	fmt.Println("")
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
