package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	originalBooks := []Book{
		{"Go in Action", "William Kennedy"},
		{"The Go Programming Language", "Alan Donovan"},
		{"Introducing Go", "Caleb Doxsey"},
	}

	fmt.Println("❌ Buggy Version:")
	var buggyPointers []*Book
	for _, book := range originalBooks {
		buggyPointers = append(buggyPointers, &book)
	}
	for _, bp := range buggyPointers {
		fmt.Printf("Title: %-30s | Address: %p\n", bp.Title, bp)
	}

	// fmt.Println("\n✅ Fixed Version:")
	// var fixedPointers []*Book
	// for i := range originalBooks {
	// 	fixedPointers = append(fixedPointers, &originalBooks[i])
	// }
	// for _, bp := range fixedPointers {
	// 	fmt.Printf("Title: %-30s | Address: %p\n", bp.Title, bp)
	// }
}
