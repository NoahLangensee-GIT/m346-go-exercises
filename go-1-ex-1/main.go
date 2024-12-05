package main

import "fmt"

func main() {

	firstName := "Noah"
	lastName := "Langensee"
	dayOfBirth, monthOfBirth, yearOfBirth := 25, 8, 2001
	numberOfSiblings := 0
	heightInMeters := 1.84
	zodiacSign := '\u264D'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
