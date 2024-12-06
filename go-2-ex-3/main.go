package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "OOP",
		117: "Codierung",
		346: "Cloud",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104)
	modules[162] = "Datenbank"
	modules[117] = "Kompression"
	fmt.Println(modules)
}
