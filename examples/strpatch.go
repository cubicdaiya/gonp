package main

import (
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
		log.Fatal("arg1 contains invalid rune")
	}

	if !utf8.ValidString(os.Args[2]) {
		log.Fatal("arg2 contains invalid rune")
	}
	a := []rune(os.Args[1])
	b := []rune(os.Args[2])
	diff := gonp.New(a, b)
	diff.Compose()

	patchedSeq := diff.Patch(a)
	fmt.Printf("success:%v, applying SES between '%s' and '%s' to '%s' is '%s'\n",
		string(b) == string(patchedSeq),
		string(a), string(b),
		string(a), string(patchedSeq))

	uniPatchedSeq, err := diff.UniPatch(a, diff.UnifiedHunks())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success:%v, applying unified format difference between '%s' and '%s' to '%s' is '%s'\n",
		string(b) == string(uniPatchedSeq),
		string(a), string(b),
		string(a), string(uniPatchedSeq))
}
