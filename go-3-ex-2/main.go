package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	switch zodiacSign {
	case Aries:
		fmt.Println("21.03 - 20.04")
	case Taurus:
		fmt.Println("21.04 - 21.05")
	case Gemini:
		fmt.Println("22.05 - 21.06")
	case Cancer:
		fmt.Println("22.06 - 22.07")
	case Leo:
		fmt.Println("23.07 - 22.08")
	case Virgo:
		fmt.Println("23.08 - 22.09")
	case Libra:
		fmt.Println("23.09 - 22.10")
	case Scorpius:
		fmt.Println("23.10 - 22.11")
	case Sagittarius:
		fmt.Println("22.11 - 20.12")
	case Capricornus:
		fmt.Println("21.12 - 19.01")
	case Aquarius:
		fmt.Println("20.01 - 18.02")
	case Pisces:
		fmt.Println("19.02 - 20.03")
	default:
		fmt.Println("not a zodiac sign")
	}

}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
	outputDateRange('\u3754')
}
