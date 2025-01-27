package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion LongestAbc.
PUNKTE: 10
BEWERTUNG:
*/

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	store := ""
	for _, x := range list {
		if len(x) >= 3 && x[:3] == "abc" {
			if len(x) > len(store) {
				store = x
			}
		}
	}
	return store
}
