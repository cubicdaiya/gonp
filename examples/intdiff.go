package main

import (
	"fmt"

	"github.com/cubicdaiya/gonp"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 9, 4, 5}
	diff := gonp.New(a, b)
	diff.Compose()
	fmt.Printf("diff %v %v\n", a, b)
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %v\n", diff.Lcs())
	fmt.Println("SES:")
	diff.PrintSes()
}
