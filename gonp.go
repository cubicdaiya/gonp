// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"math"
	"container/list"
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
	reverse bool
	path []int
	pathposi map[int]Point
	ed   int
	lcs  *list.List
	ses  *list.List
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func New(a string, b string) *Diff {
	m, n := len(a), len(b)
	if m > n {
		return &Diff{A:b, B:a, m:n, n:m, reverse:true}
	}
	return &Diff{A:a, B:b, m:m, n:n, reverse:false}
}

func (diff *Diff) Editdistance() int {
	return diff.ed
}

func (diff *Diff) Lcs() *list.List {
	return diff.lcs
}

func (diff *Diff) Ses() *list.List {
	return diff.ses
}

func (diff *Diff) Compose() {
	offset := diff.m + 1
	delta := diff.n - diff.m
	size := diff.m + diff.n + 3
	fp := make([]int, size)
	diff.path = make([]int, size)
	diff.pathposi = make(map[int]Point)
	diff.lcs = list.New()
	diff.ses = list.New()

	for i := range fp {
		fp[i] = -1
		diff.path[i] = -1
	}

	for p := 0; ; p++ {

		done := make(chan bool)
		go func() {
			for k := -p; k <= delta-1; k++ {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset)
			}
			done <- true
		}()

		go func() {
			for k := delta + p; k >= delta+1; k-- {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset], offset)
			}
			done <- true
		}()

		for i := 0; i < 2; i++ {
			<-done
		}

		fp[delta+offset] = diff.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset], offset)

		if fp[delta+offset] >= diff.n {
			diff.ed = delta + 2*p
			break
		}
	}

	r := diff.path[delta+offset]
	epc := make(map[int]Point)
	for r != -1 {
		epc[len(epc)] = Point{x:diff.pathposi[r].x, y:diff.pathposi[r].y, k:-1}
		r = diff.pathposi[r].k
	}
	diff.recordSeq(epc)
}

func (diff *Diff) recordSeq(epc map[int]Point) {
	x_idx,  y_idx  := 1, 1
	px_idx, py_idx := 0, 0
	for i:=len(epc)-1;i!=0;i-- {
		for (px_idx < epc[i].x) || (py_idx < epc[i].y) {
			var t SesType
			if (epc[i].y - epc[i].x) > (py_idx - px_idx) {
				elem := diff.B[py_idx]
				if diff.reverse {
					t = Delete
				} else {
					t = Add
				}
				diff.ses.PushBack(SesElem{c:elem, t:t})
				y_idx++
				py_idx++
			} else if epc[i].y - epc[i].x < py_idx - px_idx {
				elem := diff.A[px_idx]
				if diff.reverse {
					t = Add
				} else {
					t = Delete
				}
				diff.ses.PushBack(SesElem{c:elem, t:t})
				x_idx++
				px_idx++
			} else {
				elem := diff.A[x_idx]
				t = Common
				diff.lcs.PushBack(elem)
				diff.ses.PushBack(SesElem{c:elem, t:t})
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
		r = diff.path[k-1+offset];
	} else {
		r = diff.path[k+1+offset];
	}

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.A[x] == diff.B[y] {
		x += 1
		y += 1
	}

	diff.path[k+offset] = len(diff.pathposi)
	diff.pathposi[len(diff.pathposi)] = Point{x:x, y:y, k:r}
 
	return y
}
