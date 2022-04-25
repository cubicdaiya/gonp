package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cubicdaiya/gonp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./unifilediff filename1 filename2")
	}

	f1 := os.Args[1]
	f2 := os.Args[2]

	var (
		a   []string
		b   []string
		err error
	)

	a, err = getLines(f1)
	if err != nil {
		log.Fatalf("%s: %s", f1, err)
	}

	b, err = getLines(f2)
	if err != nil {
		log.Fatalf("%s: %s", f2, err)
	}

	th, err := buildTargetHeader(f1, f2)
	if err != nil {
		log.Fatal(err)
	}

	diff := gonp.New(a, b)
	diff.Compose()

	fmt.Printf(th.String())
	diff.PrintUniHunks(diff.UnifiedHunks())
}
