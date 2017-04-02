// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"container/list"
	"fmt"
	"math"
	"unicode/utf8"
)

type SesType int

const (
	Delete SesType = iota
	Common
	Add
)

type Point struct {
	x, y, k int
}

type SesElem struct {
	c rune
	t SesType
}

type Diff struct {
	A    []rune
	B    []rune
	m, n int
	ed   int
	ctl  *Ctl
	lcs  *list.List
	ses  *list.List
}

type Ctl struct {
	reverse  bool
	path     []int
	onlyEd   bool
	pathposi map[int]Point
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func New(a string, b string) *Diff {
	m, n := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	diff := new(Diff)
	ctl := new(Ctl)
	if m >= n {
		diff.A, diff.B = []rune(b), []rune(a)
		diff.m, diff.n = n, m
		ctl.reverse = true
	} else {
		diff.A, diff.B = []rune(a), []rune(b)
		diff.m, diff.n = m, n
		ctl.reverse = false
	}
	ctl.onlyEd = false
	diff.ctl = ctl
	return diff
}

func (diff *Diff) OnlyEd() {
	diff.ctl.onlyEd = true
}

func (diff *Diff) Editdistance() int {
	return diff.ed
}

func (diff *Diff) Lcs() string {
	var b = make([]rune, diff.lcs.Len())
	for i, e := 0, diff.lcs.Front(); e != nil; i, e = i+1, e.Next() {
		b[i] = e.Value.(rune)
	}
	return string(b)
}

func (diff *Diff) Ses() []SesElem {
	seq := make([]SesElem, diff.ses.Len())
	for i, e := 0, diff.ses.Front(); e != nil; i, e = i+1, e.Next() {
		seq[i].c = e.Value.(SesElem).c
		seq[i].t = e.Value.(SesElem).t
	}
	return seq
}

func (diff *Diff) PrintSes() {
	for _, e := 0, diff.ses.Front(); e != nil; e = e.Next() {
		ee := e.Value.(SesElem)
		switch ee.t {
		case Delete:
			fmt.Println("- " + string(ee.c))
		case Add:
			fmt.Println("+ " + string(ee.c))
		case Common:
			fmt.Println("  " + string(ee.c))
		}
	}
}

func (diff *Diff) Compose() {
	offset := diff.m + 1
	delta := diff.n - diff.m
	size := diff.m + diff.n + 3
	fp := make([]int, size)
	diff.ctl.path = make([]int, size)
	diff.ctl.pathposi = make(map[int]Point)
	diff.lcs = list.New()
	diff.ses = list.New()

	for i := range fp {
		fp[i] = -1
		diff.ctl.path[i] = -1
	}

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

	if diff.ctl.onlyEd {
		return
	}

	r := diff.ctl.path[delta+offset]
	epc := make(map[int]Point)
	for r != -1 {
		epc[len(epc)] = Point{x: diff.ctl.pathposi[r].x, y: diff.ctl.pathposi[r].y, k: -1}
		r = diff.ctl.pathposi[r].k
	}
	diff.recordSeq(epc)
}

func (diff *Diff) recordSeq(epc map[int]Point) {
	x_idx, y_idx := 1, 1
	px_idx, py_idx := 0, 0
	ctl := diff.ctl
	for i := len(epc) - 1; i >= 0; i-- {
		for (px_idx < epc[i].x) || (py_idx < epc[i].y) {
			var t SesType
			if (epc[i].y - epc[i].x) > (py_idx - px_idx) {
				elem := diff.B[py_idx]
				if ctl.reverse {
					t = Delete
				} else {
					t = Add
				}
				diff.ses.PushBack(SesElem{c: elem, t: t})
				y_idx++
				py_idx++
			} else if epc[i].y-epc[i].x < py_idx-px_idx {
				elem := diff.A[px_idx]
				if ctl.reverse {
					t = Add
				} else {
					t = Delete
				}
				diff.ses.PushBack(SesElem{c: elem, t: t})
				x_idx++
				px_idx++
			} else {
				elem := diff.A[px_idx]
				t = Common
				diff.lcs.PushBack(elem)
				diff.ses.PushBack(SesElem{c: elem, t: t})
				x_idx++
				y_idx++
				px_idx++
				py_idx++
			}
		}
	}
}

func (diff *Diff) snake(k, p, pp, offset int) int {
	r := 0
	if p > pp {
		r = diff.ctl.path[k-1+offset]
	} else {
		r = diff.ctl.path[k+1+offset]
	}

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.A[x] == diff.B[y] {
		x++
		y++
	}

	if !diff.ctl.onlyEd {
		diff.ctl.path[k+offset] = len(diff.ctl.pathposi)
		diff.ctl.pathposi[len(diff.ctl.pathposi)] = Point{x: x, y: y, k: r}
	}

	return y
}
