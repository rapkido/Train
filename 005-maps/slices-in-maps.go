package main

import (
	"fmt"
)

func main() {
	person := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"moneypenny_miss", "James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}
	person["coskuner_bestami"] = []string{"Bilgisayar", "Go", "insanlar", "Python"}
	delete(person, "coskuner_bestami")
	for i, per := range person {
		fmt.Printf("Kayıt: %v\n", i)
		for j, data := range per {
			fmt.Printf("\t%v. içerik: %v\n", j, data)
		}

	}
}
