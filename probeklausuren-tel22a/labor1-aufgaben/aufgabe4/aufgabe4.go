package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Power2.
Die Funktion muss rekursiv sein.
PUNKTE: 5
BEWERTUNG:
*/

// Power2 erwartet einen int-Parameter x und liefert die Potenz "2 hoch x".
func Power2(x int) int {

	// 2^0 == 1
	if x == 0 {
		return 1
	}

	// 2^x == 2 * 2^(x-1)
	return 2 * Power2(x-1)
}
