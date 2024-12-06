package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   int
	MonthOfBirth int
	YearOfBirth  int
}

type Profile struct {
	Name             FullName
	Born             BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Noah",
			LastName:  "Langensee",
		},
		Born: BirthDate{
			DayOfBirth:   25,
			MonthOfBirth: 8,
			YearOfBirth:  2001,
		},
		NumberOfSiblings: 0,
		ZodiacSign:       '\u264D',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
