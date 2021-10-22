package main

import "fmt"

func main() {
	text := "Hello World"
	fmt.Println(text)

	// Aufruf der Funktion plus2
	plus2(77) // Ergebnis wird nicht verwendet, Aufruf hat keinen Effekt.

	// Richtig ist:
	x := plus2(3) // Funktion aufrufen und Ergebnis in x speichern
	fmt.Println(x)

	testPlus2()

	//  fmt.Printf("Ergebnis: %v", x)
}

// Definition einer Funktion
// Name: "plus2"
// Parameter: Eine Zahl x vom Typ int
// - wird beim Aufruf der Funktion eingesetzt
// Rückgabetyp: Eine Zahl
func plus2(x int) int { // "Signatur"
	result := x + 2 + 1
	fmt.Println("func plus2 wurde ausgeführt")
	return result // "return": Die Funktion ist zu Ende, das Ergebnis ist x+2
}

func checkPlus2(expected int, actual int) {
	if actual != expected {
		fmt.Printf("Fehler, plus2(3) sollte %v liefern, tut es aber nicht.\n", expected)
	}
}

func testPlus2() {
	checkPlus2(5, plus2(3))
	checkPlus2(-13, plus2(-15))

}

// "Fehler, plus(40) sollte 42 liefern, es wurde aber 43 geliefert."
