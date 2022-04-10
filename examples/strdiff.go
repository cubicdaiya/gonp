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
		log.Fatal("./strdiff arg1 arg2")
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
	fmt.Printf("Editdistance: %d\n", diff.Editdistance())
	fmt.Printf("LCS: %s\n", string(diff.Lcs()))
	fmt.Println("SES:")

	var buf bytes.Buffer
	ses := diff.Ses()
	for _, e := range ses {
		ee := e.GetElem()
		switch e.GetType() {
		case gonp.SesDelete:
			fmt.Fprintf(&buf, "-%c\n", ee)
		case gonp.SesAdd:
			fmt.Fprintf(&buf, "+%c\n", ee)
		case gonp.SesCommon:
			fmt.Fprintf(&buf, " %c\n", ee)
		}
	}
	fmt.Print(buf.String())
}
