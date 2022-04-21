package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cubicdaiya/gonp"
)

// Target consists of a path and mtime of file.
type Target struct {
	fname string
	mtime time.Time
}

// TargetHeader has 2 targets based on pathes and mtimes based on 2 files
type TargetHeader struct {
	targets []Target
}

// getLines returns a file contents as string array
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

// builderTargetHeader returns TargetHeader constructed based on 2 files given as arguments
func buildTargetHeader(f1, f2 string) (TargetHeader, error) {
	fi1, err := os.Stat(f1)
	if err != nil {
		return TargetHeader{}, err
	}
	fi2, err := os.Stat(f2)
	if err != nil {
		return TargetHeader{}, err
	}
	return TargetHeader{
		targets: []Target{
			Target{fname: f1, mtime: fi1.ModTime()},
			Target{fname: f2, mtime: fi2.ModTime()},
		},
	}, nil
}

// String returns a content of TargetHeader as a string
func (th *TargetHeader) String() string {
	if len(th.targets) != 2 {
		return ""
	}
	var b bytes.Buffer
	fmt.Fprintf(&b, "--- %s\t%s\n", th.targets[0].fname, th.targets[0].mtime.Format(time.RFC3339Nano))
	fmt.Fprintf(&b, "+++ %s\t%s\n", th.targets[1].fname, th.targets[1].mtime.Format(time.RFC3339Nano))
	return b.String()
}

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
