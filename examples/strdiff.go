package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/cubicdaiya/gonp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./strdiff arg1 arg2")
	}
	a := []rune(os.Args[1])
	b := []rune(os.Args[2])
	diff := gonp.New[rune](a, b)
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
			fmt.Fprintf(&buf, "- %c\n", ee)
		case gonp.SesAdd:
			fmt.Fprintf(&buf, "+ %c\n", ee)
		case gonp.SesCommon:
			fmt.Fprintf(&buf, "  %c\n", ee)
		}
	}
	fmt.Print(buf.String())
}
