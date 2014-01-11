// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"container/list"
	"math"
)

type SesType int

const (
	Delete SesType = 0
	Common SesType = 1
	Add    SesType = 2
)

type Point struct {
	x, y, k int
}

type SesElem struct {
	c byte
	t SesType
}

type Diff struct {
	A    string
	B    string
	m, n int
	ed   int
	ctl  *Ctl
	lcs  *list.List
	ses  *list.List
}

type Ctl struct {
	reverse  bool
	path     []int
	pathposi map[int]Point
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func New(a string, b string) *Diff {
	m, n := len(a), len(b)
	diff := new(Diff)
	ctl := new(Ctl)
	if m > n {
		diff.A, diff.B = b, a
		diff.m, diff.n = n, m
		ctl.reverse = true
	} else {
		diff.A, diff.B = a, b
		diff.m, diff.n = m, n
		ctl.reverse = false
	}
	diff.ctl = ctl
	return diff
}

func (diff *Diff) Editdistance() int {
	return diff.ed
}

func (diff *Diff) Lcs() string {
	var b = make([]byte, diff.lcs.Len())
	for i, e := 0, diff.lcs.Front(); e != nil; i, e = i+1, e.Next() {
		b[i] = e.Value.(byte)
	}
	return string(b)
}

func (diff *Diff) Ses() *list.List {
	return diff.ses
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
	ctl := diff.ctl

	for i := range fp {
		fp[i] = -1
		ctl.path[i] = -1
	}

	for p := 0; ; p++ {

		for k := -p; k <= delta-1; k++ {
			fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset, diff.ctl)
		}

		for k := delta + p; k >= delta+1; k-- {
			fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset, diff.ctl)
		}

		fp[delta+offset] = diff.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset], offset, diff.ctl)

		if fp[delta+offset] >= diff.n {
			diff.ed = delta + 2*p
			break
		}
	}

	r := ctl.path[delta+offset]
	epc := make(map[int]Point)
	for r != -1 {
		epc[len(epc)] = Point{x: ctl.pathposi[r].x, y: ctl.pathposi[r].y, k: -1}
		r = ctl.pathposi[r].k
	}
	diff.recordSeq(epc)
}

func (diff *Diff) recordSeq(epc map[int]Point) {
	x_idx, y_idx := 1, 1
	px_idx, py_idx := 0, 0
	ctl := diff.ctl
	for i := len(epc) - 1; i != 0; i-- {
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

func (diff *Diff) snake(k, p, pp, offset int, ctl *Ctl) int {
	r := 0
	if p > pp {
		r = ctl.path[k-1+offset]
	} else {
		r = ctl.path[k+1+offset]
	}

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.A[x] == diff.B[y] {
		x += 1
		y += 1
	}

	ctl.path[k+offset] = len(ctl.pathposi)
	ctl.pathposi[len(ctl.pathposi)] = Point{x: x, y: y, k: r}

	return y
}
