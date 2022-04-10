package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/cubicdaiya/gonp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./unistrdiff arg1 arg2")
	}
	if !utf8.ValidString(os.Args[1]) {
		log.Fatalf("arg1 contains invalid rune")
	}

	if !utf8.ValidString(os.Args[2]) {
		log.Fatalf("arg2 contains invalid rune")
	}
	a := []rune(os.Args[1])
	b := []rune(os.Args[2])
	diff := gonp.New(a, b)
	diff.Compose()
	fmt.Printf("Editdistance:%d\n", diff.Editdistance())
	fmt.Printf("LCS:%s\n", string(diff.Lcs()))
	//diff.PrintUniHunks(diff.UnifiedHunks())

	fmt.Println("Unified format difference:")
	uniHunks := diff.UnifiedHunks()
	var w bytes.Buffer
	for _, uniHunk := range uniHunks {
		fmt.Fprintf(&w, uniHunk.SprintDiffRange())
		for _, e := range uniHunk.GetChanges() {
			switch e.GetType() {
			case gonp.SesDelete:
				fmt.Fprintf(&w, "-%c\n", e.GetElem())
			case gonp.SesAdd:
				fmt.Fprintf(&w, "+%c\n", e.GetElem())
			case gonp.SesCommon:
				fmt.Fprintf(&w, " %c\n", e.GetElem())
			}
		}
	}
	fmt.Print(w.String())

}
