package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ArraySums.
PUNKTE: 15
BEWERTUNG:
*/

// ArraySums erwartet eine int-Slice l erwartet und liefert eine int-Slice,
// die an Stelle n die Summe der Elemente aus l bis zu Stelle n enthält.
func ArraySums(l []int) []int {
	result := []int{}
	sum := 0
	for _, v := range l {
		sum = sum + v
		result = append(result, sum)
	}
	return result
}
