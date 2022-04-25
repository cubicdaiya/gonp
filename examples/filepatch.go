package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cubicdaiya/gonp"
)

func equalsStringSlice(a, b []string) bool {
	m, n := len(a), len(b)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./filepatch filename1 filename2")
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

	diff := gonp.New(a, b)
	diff.Compose()

	patchedSeq := diff.Patch(a)
	fmt.Printf("success:%v, applying SES between '%s' and '%s'\n", equalsStringSlice(b, patchedSeq), f1, f2)

	uniPatchedSeq, err := diff.UniPatch(a, diff.UnifiedHunks())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success:%v, applying unified format difference between '%s' and '%s'\n", equalsStringSlice(b, uniPatchedSeq), f1, f2)
}
