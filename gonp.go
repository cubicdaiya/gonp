// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

const (
	// SesDelete is manipulaton type of deleting element in SES
	SesDelete SesType = iota
	// SesCommon is manipulaton type of same element in SES
	SesCommon
	// SesAdd is manipulaton type of adding element in SES
	SesAdd
)

// SesType is manipulaton type
type SesType int

// Point is coordinate in edit graph
type Point struct {
	x, y int
}

// PointWithRoute is coordinate in edit graph attached route
type PointWithRoute struct {
	x, y, r int
}

// SesElem is element of SES
type SesElem struct {
	e rune
	t SesType
}

// Diff is context for calculating difference between a and b
type Diff struct {
	a              []rune
	b              []rune
	m, n           int
	ed             int
	lcs            []rune
	ses            []SesElem
	reverse        bool
	path           []int
	onlyEd         bool
	pointWithRoute []PointWithRoute
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// New is initializer of Diff
func New(a, b string) *Diff {
	m, n := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	diff := new(Diff)
	diff.a, diff.b = []rune(a), []rune(b)
	diff.m, diff.n = m, n
	diff.reverse = false
	if m >= n {
		diff.a, diff.b = diff.b, diff.a
		diff.m, diff.n = n, m
		diff.reverse = true
	}
	diff.onlyEd = false
	return diff
}

// OnlyEd enables to calculate only edit distance
func (diff *Diff) OnlyEd() {
	diff.onlyEd = true
}

// Editdistance returns edit distance between a and b
func (diff *Diff) Editdistance() int {
	return diff.ed
}

// Lcs returns LCS (Longest Common Subsequence) between a and b
func (diff *Diff) Lcs() []rune {
	return diff.lcs
}

// Lcs returns LCS (Longest Common Subsequence) string between a and b
func (diff *Diff) LcsString() string {
	return string(diff.lcs)
}

// Ses return SES (Shortest Edit Script) between a and b
func (diff *Diff) Ses() []SesElem {
	return diff.ses
}

// PrintSes prints shortest edit script between a and b
func (diff *Diff) PrintSes() {
	fmt.Print(diff.SprintSes())
}

// SprintSes returns string about shortest edit script between a and b
func (diff *Diff) SprintSes() string {
	var buf bytes.Buffer
	diff.FprintSes(&buf)
	return buf.String()
}

// FprintSes emit about shortest edit script between a and b to w
func (diff *Diff) FprintSes(w io.Writer) {
	for _, e := range diff.ses {
		switch e.t {
		case SesDelete:
			fmt.Fprintf(w, "- %c\n", e.e)
		case SesAdd:
			fmt.Fprintf(w, "+ %c\n", e.e)
		case SesCommon:
			fmt.Fprintf(w, "  %c\n", e.e)
		}
	}
}

// Compose composes diff between a and b
func (diff *Diff) Compose() {
	fp := make([]int, diff.m+diff.n+3)
	diff.path = make([]int, diff.m+diff.n+3)
	diff.pointWithRoute = make([]PointWithRoute, 0)

	for i := range fp {
		fp[i] = -1
		diff.path[i] = -1
	}

	offset := diff.m + 1
	delta := diff.n - diff.m
	for p := 0; ; p++ {

		for k := -p; k <= delta-1; k++ {
			fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset)
		}

		for k := delta + p; k >= delta+1; k-- {
			fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset)
		}

		fp[delta+offset] = diff.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset], offset)

		if fp[delta+offset] >= diff.n {
			diff.ed = delta + 2*p
			break
		}
	}

	if diff.onlyEd {
		return
	}

	r := diff.path[delta+offset]
	epc := make([]Point, 0)
	for r != -1 {
		epc = append(epc, Point{x: diff.pointWithRoute[r].x, y: diff.pointWithRoute[r].y})
		r = diff.pointWithRoute[r].r
	}
	diff.recordSeq(epc)
}

func (diff *Diff) snake(k, p, pp, offset int) int {
	r := 0
	if p > pp {
		r = diff.path[k-1+offset]
	} else {
		r = diff.path[k+1+offset]
	}

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.a[x] == diff.b[y] {
		x++
		y++
	}

	if !diff.onlyEd {
		diff.path[k+offset] = len(diff.pointWithRoute)
		diff.pointWithRoute = append(diff.pointWithRoute, PointWithRoute{x: x, y: y, r: r})
	}

	return y
}

func (diff *Diff) recordSeq(epc []Point) {
	x, y := 1, 1
	px, py := 0, 0
	for i := len(epc) - 1; i >= 0; i-- {
		for (px < epc[i].x) || (py < epc[i].y) {
			if (epc[i].y - epc[i].x) > (py - px) {
				t := SesAdd
				if diff.reverse {
					t = SesDelete
				}
				diff.ses = append(diff.ses, SesElem{e: diff.b[py], t: t})
				y++
				py++
			} else if epc[i].y-epc[i].x < py-px {
				t := SesDelete
				if diff.reverse {
					t = SesAdd
				}
				diff.ses = append(diff.ses, SesElem{e: diff.a[px], t: t})
				x++
				px++
			} else {
				diff.lcs = append(diff.lcs, diff.a[px])
				diff.ses = append(diff.ses, SesElem{e: diff.a[px], t: SesCommon})
				x++
				y++
				px++
				py++
			}
		}
	}
}
