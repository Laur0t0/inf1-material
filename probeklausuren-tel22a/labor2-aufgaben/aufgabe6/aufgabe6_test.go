package aufgabe6

import "fmt"

func ExampleCommonFriends() {
	f1 := Person{"Gödel", []Person{}}
	f2 := Person{"Turing", []Person{}}
	f3 := Person{"Euler", []Person{}}
	f4 := Person{"Gauss", []Person{}}
	f5 := Person{"Lovelace", []Person{}}

	p1 := Person{"Lamport", []Person{f1, f3, f5}}
	p2 := Person{"Knuth", []Person{f2, f3, f4}}
	p3 := Person{"Finn", []Person{f4, f3, f2}}
	p4 := Person{"Fonner", []Person{f4, f3, f2, f1}}

	fmt.Println(CommonFriends(p1, p2))
	fmt.Println(CommonFriends(p1, p1))
	fmt.Println(CommonFriends(p2, p2))
	fmt.Println(CommonFriends(p2, p3))
	fmt.Println(CommonFriends(p4, p4))

	// Output:
	// [{Euler []}]
	// [{Gödel []} {Euler []} {Lovelace []}]
	// [{Turing []} {Euler []} {Gauss []}]
	// [{Turing []} {Euler []} {Gauss []}]
	// [{Gauss []} {Euler []} {Turing []} {Gödel []}]

}
