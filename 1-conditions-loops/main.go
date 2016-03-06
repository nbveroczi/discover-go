package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(100)
	school := "Holberton School"
	var beautifulWeather = true

	if randomNumber > 50 {
		answer := "is"
		fmt.Printf("my random number is %d and %s greater than 50\n", randomNumber, answer)
	} else {
		answer := "is not"
		fmt.Printf("my random number is %d and %s greater than 50\n", randomNumber, answer)
	}

	if school == "Holberton School" {
		fmt.Printf("I am a student of %s\n", school)
	} else {
		fmt.Printf("I am NOT a student of %s\n", school)
	}

	if beautifulWeather == true {
		fmt.Printf("It's beautiful weather!\n")
	} else {
		fmt.Printf("The weather is miserable!\n")
	}

	for _, holbertonFounders := range [3]string{"Rudy", "Julien", "Sylvain"} {
		fmt.Printf("%v is a founder\n", holbertonFounders)
	}
}
