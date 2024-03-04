package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Difference.
PUNKTE: 10
BEWERTUNG:
*/

// Difference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen aus l1,
// die nicht in l2 vorhanden sind.
func Difference(l1, l2 []int) []int {
	result := []int{}

	for _, x := range l1 {
		same := false
		for _, y := range l2 {
			if x == y {
				same = true
			}
		}
		if !same {
			result = append(result, x)
		}
	}
	return result
}

// 	for i := 0; i < len(l1); i++ {
// 		for n := 0; n < len(l1); n++ {
// 			if l1[i] == same(l1, l2) {
// 				result = append(result, l1[i])
// 			}
// 		}
// 	}
// 	return result
// }

// func same(l1, l2 []int) int {
// 	for i := 0; i < len(l1); i++ {
// 		for n := 0; n < len(l2); n++ {
// 			if l1[i] == l2[n] {
// 				return l1[i]
// 			}
// 		}
// 	}
// 	return 0
// }
