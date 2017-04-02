package main

import (
	"fmt"
	"github.com/cubicdaiya/gonp"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("./strdiff arg1 arg2")
		os.Exit(1)
	}
	diff := gonp.New(os.Args[1], os.Args[2])
	diff.Compose()
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", diff.Lcs())
	fmt.Println("SES:")
	diff.PrintSes()
}
