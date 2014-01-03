// The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
// by described by Sun Wu, Udi Manber and Gene Myers

package gonp

import (
	"math"
)

type Diff struct {
	A    string
	B    string
	m, n int
	ed   int
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func New(a string, b string) *Diff {
	m, n := len(a), len(b)
	if m > n {
		return &Diff{A:b, B:a, m:n, n:m}
	}
	return &Diff{A:a, B:b, m:m, n:n}
}

func (diff *Diff) Editdistance() int {
	return diff.ed
}

func (diff *Diff) Compose() {

	offset := diff.m + 1
	delta := diff.n - diff.m
	size := diff.m + diff.n + 3
	fp := make([]int, size)

	for i := range fp {
		fp[i] = -1
	}

	for p := 0; ; p++ {

		done := make(chan bool)
		go func() {
			for k := -p; k <= delta-1; k++ {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}

			done <- true
		}()

		go func() {
			for k := delta + p; k >= delta+1; k-- {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}
			done <- true
		}()

		for i := 0; i < 2; i++ {
			<-done
		}

		fp[delta+offset] = diff.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset])

		if fp[delta+offset] >= diff.n {
			diff.ed = delta + 2*p
			return
		}
	}
}

func (diff *Diff) snake(k, p, pp int) int {

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.A[x] == diff.B[y] {
		x += 1
		y += 1
	}

	return y
}
