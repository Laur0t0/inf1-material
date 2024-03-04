package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Person vorgegeben.
Vervollständigen Sie die Funktion CommonFriends.
PUNKTE: 15
BEWERTUNG:
*/

// Person repräsentiert eine Person in einem sozialen Netzwerk.
type Person struct {
	Name    string
	Friends []Person
}

// CommonFriends erwartet zwei Personen und liefert eine Liste
// mit allen Personen, die mit beiden befreundet sind.
func CommonFriends(per1, per2 Person) []Person {
	result := []Person{}
	for i := 0; i < len(per1.Friends); i++ {
		for n := 0; n < len(per2.Friends); n++ {
			if per1.Friends[i].Name == per2.Friends[n].Name {
				result = append(result, per1.Friends[i])
			}
		}
	}
	return result
}
