package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}
	type Class struct {
		Students []Student
		Modules  map[int]string
	}
	Ina23al := Class{
		Students: []Student{
			{
				FirstName: "John",
				LastName:  "Doe",
			},
			{
				FirstName: "Jane",
				LastName:  "Jefferson",
			},
			{
				FirstName: "Jenny",
				LastName:  "Lopez",
			},
		},
		Modules: map[int]string{
			104: "OOP",
			117: "Codierung",
			346: "Cloud",
		},
	}
	Ina24al := Class{
		Students: []Student{
			{
				FirstName: "Peter",
				LastName:  "Parker",
			},
			{
				FirstName: "Hans",
				LastName:  "Zimmerli",
			},
			{
				FirstName: "Mario",
				LastName:  "Müller",
			},
		},
		Modules: map[int]string{
			117: "Codierung",
			346: "Cloud",
		},
	}
	classes := []Class{Ina23al, Ina24al}
	fmt.Println(classes)
	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()
}
