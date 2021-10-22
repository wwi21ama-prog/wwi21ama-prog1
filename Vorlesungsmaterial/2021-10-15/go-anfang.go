// Paketangabe, meistens main (jetzt am Anfang)
package main

// Wir geben an, dass wir das Paket "fmt" brauchen.
import "fmt"

// main()-Funktion, Entry point des Programms
func main() {

	// a := 42
	// var b int = 77

	p1 := 1 * 2 * 3 * 4 * 5
	fmt.Println(p1)

	fmt.Println(factorial(5))
	fmt.Println(factorial2(5))

	//fmt.Println(produkt)

	fmt.Println(sumMultiples(3, 10))
	fmt.Println(sumMultiples(5, 27))

	testSumMultiples()

}

//liefert die Fakultät von n
func factorial(n int) int {
	produkt := 1
	for zaehler := 1; zaehler <= n; zaehler++ {
		produkt = produkt * zaehler
	}
	return produkt
}

//Liefert auch die Fakultät von n, rekursive Version

func factorial2(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial2(n-1)
}

func sumMultiples(x, n int) int { // Beispiel: x = 5
	result := 0
	vielfaches := x // xOld = 5
	for vielfaches < n {
		result += vielfaches
		vielfaches += x // x = x + 5
	}

	// TODO: Hier das Ergebnis berechnen.
	return result
}

func testSumMultiples() {
	expectEqual(sumMultiples(3, 10), 18)
	expectEqual(sumMultiples(2, 10), 20)
	expectEqual(sumMultiples(10, 2), 0)
}

// Hilfsfunktion, wird in den Tests benutzt
func expectEqual(value, expected interface{}) {
	if value != expected {
		fmt.Printf("Fehler: %v bekommen, erwartet war aber %v.\n", value, expected)
	} else {
		fmt.Printf("OK: %v bekommen, erwartet war aber %v.\n", value, expected)
	}
}
