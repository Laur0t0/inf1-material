package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MultFive.
PUNKTE: 5
BEWERTUNG:
*/

// MultFive erwartet eine Liste von Zahlen und liefert
// das Produkt der Elemente an den durch 5 teilbaren Positionen.
func MultFive(list []int) int {
	result := 1
	if len(list) == 0 {
		return 1
	}
	for v, x := range list {
		if v%5 == 0 {
			result = result * x
		}
	}
	return result
}
