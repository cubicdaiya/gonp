package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cubicdaiya/gonp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./strdiff arg1 arg2")
	}
	diff := gonp.New(os.Args[1], os.Args[2])
	diff.Compose()
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", diff.Lcs())
	fmt.Println("SES:")
	diff.PrintSes()
}
