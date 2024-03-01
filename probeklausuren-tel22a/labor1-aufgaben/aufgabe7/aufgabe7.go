package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Intersect.
PUNKTE: 10
BEWERTUNG:
*/

// Intersect erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen, die sowohl in l1 als auch
// in l2 vorhanden sind.
// Hinweis: Gehen Sie davon aus, dass jedes Element pro Liste höchstens einmal vorkommt.
func Intersect(l1, l2 []int) []int {
	result := []int{}

	for _, v1 := range l1 {
		for _, v2 := range l2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}

	return result
}
