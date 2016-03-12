package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:",name"`
	DOB  string `json:",date_of_birth"`
	City string `json:",city"`
}

func main() {

	u := user{Name: "Betty HolbertonMarch", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)
	dob, err := time.Parse("January 2, 2006", u.DOB)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s who was born in %s would be %d years old today\n", u.Name, u.City, time.Now().Year()-dob.Year())
}
