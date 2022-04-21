// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"bytes"
	"fmt"
	"io"
)

const (
	// SesDelete is manipulaton type of deleting element in SES
	SesDelete SesType = iota
	// SesCommon is manipulaton type of same element in SES
	SesCommon
	// SesAdd is manipulaton type of adding element in SES
	SesAdd
)

const (
	// limit of cordinate size
	DefaultRouteSize = 2000000
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

// Type constraints for element in SES
type Elem interface {
	// int32 overlaps rune
	~rune | ~string | ~byte | ~int | ~int8 | ~int16 | ~int64 | ~float32 | ~float64
}

// SesElem is element of SES
type SesElem[T Elem] struct {
	e    T
	t    SesType
	aIdx int
	bIdx int
}

// Diff is context for calculating difference between a and b
type Diff[T Elem] struct {
	a, b           []T
	m, n           int
	ox, oy         int
	ed             int
	lcs            []T
	ses            []SesElem[T]
	reverse        bool
	path           []int
	onlyEd         bool
	pointWithRoute []PointWithRoute
	contextSize    int
	routeSize      int
}

// New is initializer of Diff
func New[T Elem](a, b []T) *Diff[T] {
	diff := new(Diff[T])
	m, n := len(a), len(b)
	reverse := false
	if m >= n {
		a, b = b, a
		m, n = n, m
		reverse = true
	}
	diff.a, diff.b = a, b
	diff.m, diff.n = m, n
	diff.ed = 0
	diff.reverse = reverse
	diff.onlyEd = false
	diff.contextSize = DefaultContextSize
	diff.routeSize = DefaultRouteSize
	return diff
}

// OnlyEd enables to calculate only edit distance
func (diff *Diff[T]) OnlyEd() {
	diff.onlyEd = true
}

// SetContextSize sets the context size of unified format difference
func (diff *Diff[T]) SetContextSize(n int) {
	diff.contextSize = n
}

// SetRouteSize sets the context size of unified format difference
func (diff *Diff[T]) SetRouteSize(n int) {
	diff.routeSize = n
}

// GetElem is getter of element of SES
func (e *SesElem[T]) GetElem() T {
	return e.e
}

// GetType is getter of manipulation type of SES
func (e *SesElem[T]) GetType() SesType {
	return e.t
}

// Editdistance returns edit distance between a and b
func (diff *Diff[T]) Editdistance() int {
	return diff.ed
}

// Lcs returns LCS (Longest Common Subsequence) between a and b
func (diff *Diff[T]) Lcs() []T {
	return diff.lcs
}

// Ses return SES (Shortest Edit Script) between a and b
func (diff *Diff[T]) Ses() []SesElem[T] {
	return diff.ses
}

// PrintSes prints shortest edit script between a and b
func (diff *Diff[T]) PrintSes() {
	fmt.Print(diff.SprintSes())
}

// SprintSes returns string about shortest edit script between a and b
func (diff *Diff[T]) SprintSes() string {
	var buf bytes.Buffer
	diff.FprintSes(&buf)
	return buf.String()
}

// FprintSes emit about shortest edit script between a and b to w
func (diff *Diff[T]) FprintSes(w io.Writer) {
	for _, e := range diff.ses {
		switch e.t {
		case SesDelete:
			fmt.Fprintf(w, "-%v\n", e.e)
		case SesAdd:
			fmt.Fprintf(w, "+%v\n", e.e)
		case SesCommon:
			fmt.Fprintf(w, " %v\n", e.e)
		}
	}
}

// Compose composes diff between a and b
func (diff *Diff[T]) Compose() {
ONP:
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

		if fp[delta+offset] >= diff.n || len(diff.pointWithRoute) > diff.routeSize {
			diff.ed += delta + 2*p
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

	if !diff.recordSeq(epc) {
		goto ONP
	}
}

func (diff *Diff[T]) snake(k, p, pp, offset int) int {
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

func (diff *Diff[T]) recordSeq(epc []Point) bool {
	x, y := 1, 1
	px, py := 0, 0
	for i := len(epc) - 1; i >= 0; i-- {
		for (px < epc[i].x) || (py < epc[i].y) {
			if (epc[i].y - epc[i].x) > (py - px) {
				if diff.reverse {
					diff.ses = append(diff.ses, SesElem[T]{e: diff.b[py], t: SesDelete, aIdx: y + diff.oy, bIdx: 0})
				} else {
					diff.ses = append(diff.ses, SesElem[T]{e: diff.b[py], t: SesAdd, aIdx: 0, bIdx: y + diff.oy})
				}
				y++
				py++
			} else if epc[i].y-epc[i].x < py-px {
				if diff.reverse {
					diff.ses = append(diff.ses, SesElem[T]{e: diff.a[px], t: SesAdd, aIdx: 0, bIdx: x + diff.ox})
				} else {
					diff.ses = append(diff.ses, SesElem[T]{e: diff.a[px], t: SesDelete, aIdx: x + diff.ox, bIdx: 0})

				}
				x++
				px++
			} else {
				if diff.reverse {
					diff.lcs = append(diff.lcs, diff.b[py])
					diff.ses = append(diff.ses, SesElem[T]{e: diff.b[py], t: SesCommon, aIdx: y + diff.oy, bIdx: x + diff.ox})
				} else {
					diff.lcs = append(diff.lcs, diff.a[px])
					diff.ses = append(diff.ses, SesElem[T]{e: diff.a[px], t: SesCommon, aIdx: x + diff.ox, bIdx: y + diff.oy})
				}
				x++
				y++
				px++
				py++
			}
		}
	}

	if x > diff.m && y > diff.n {
		// all recording succeeded
	} else {
		diff.a = diff.a[x-1:]
		diff.b = diff.b[y-1:]
		diff.m = len(diff.a)
		diff.n = len(diff.b)
		diff.ox = x - 1
		diff.oy = y - 1
		return false
	}

	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
