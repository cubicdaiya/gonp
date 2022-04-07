package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cubicdaiya/gonp"
)

func getLines(f string) ([]string, error) {
	fp, err := os.Open(f)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("./unifilediff filename11 filename2")
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
	diff.PrintUniHunks(diff.UnifiedHunks())
}
