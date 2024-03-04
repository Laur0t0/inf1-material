package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion GetStringsBetween.
PUNKTE: 15
BEWERTUNG:
*/

// GetStringsBetween eine Liste und zwei Strings first und last.
// Die Funktion liefert die Teilliste, die zwischen first und last liegt.
// first sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func GetStringsBetween(list []string, first, last string) []string {
	result := []string{}
	temp1 := len(list)
	temp2 := len(list)
	for v, x := range list {
		if x == first {
			temp1 = v
		}
		if x == last {
			temp2 = v
		}
		if v > temp1 && v < temp2 {
			result = append(result, x)
		}
	}
	return result
}
