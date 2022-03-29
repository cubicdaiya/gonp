package gonp

import (
	"bytes"
	"fmt"
	"io"
)

const (
	PhaseFrontDiff = iota
	PhaseInDiff
	PhaseBehindDiff
)

const (
	DefaultContextSize = 3
)

type UniHunk[T Elem] struct {
	a, b, c, d int // @@ -a,b +c,d @@
	changes    []SesElem[T]
}

func (uniHunk *UniHunk[T]) GetChanges() []SesElem[T] {
	return uniHunk.changes
}

func (uniHunk *UniHunk[T]) SprintDiffRange() string {
	return fmt.Sprintf("@@ -%d,%d +%d,%d @@\n", uniHunk.a, uniHunk.b, uniHunk.c, uniHunk.d)
}

// PrintUniHunks prints unified format difference between and b
func (diff *Diff[T]) PrintUniHunks(uniHunks []UniHunk[T]) {
	fmt.Print(diff.SprintUniHunks(uniHunks))
}

// SprintUniHunks returns string about unified format difference between a and b
func (diff *Diff[T]) SprintUniHunks(uniHunks []UniHunk[T]) string {
	var buf bytes.Buffer
	diff.FprintUniHunks(&buf, uniHunks)
	return buf.String()
}

// FprintUniHunks emit about unified format difference between a and b to w
func (diff *Diff[T]) FprintUniHunks(w io.Writer, uniHunks []UniHunk[T]) {
	for _, uniHunk := range uniHunks {
		fmt.Fprintf(w, uniHunk.SprintDiffRange())
		for _, e := range uniHunk.GetChanges() {
			switch e.GetType() {
			case SesDelete:
				fmt.Fprintf(w, "-%v\n", e.GetElem())
			case SesAdd:
				fmt.Fprintf(w, "+%v\n", e.GetElem())
			case SesCommon:
				fmt.Fprintf(w, " %v\n", e.GetElem())
			}
		}
	}
}

func (diff *Diff[T]) shiftCommons(commons []SesElem[T], b, d int) ([]SesElem[T], int, int) {
	if len(commons) > diff.contextSize {
		commons = commons[1:]
		b -= 1
		d -= 1
	}
	return commons, b, d
}

// Compose composes unified format difference between a and b
func (diff *Diff[T]) UnifiedHunks() []UniHunk[T] {
	if diff.ed == 0 {
		return []UniHunk[T]{}
	}
	uniHunks := make([]UniHunk[T], 0)
	changes := make([]SesElem[T], 0)
	phase := PhaseFrontDiff
	cc := 0
	b, d := 0, 0

	for i, e := range diff.ses {
		switch e.t {
		case SesDelete:
			fallthrough
		case SesAdd:
			switch phase {
			case PhaseFrontDiff:
				phase = PhaseInDiff
				changes = append(changes, e)
			case PhaseInDiff:
				changes = append(changes, e)
				cc = 0
			case PhaseBehindDiff:
				// do nothing
			}
			if e.t == SesDelete {
				b += 1
			} else {
				d += 1
			}
		case SesCommon:
			switch phase {
			case PhaseFrontDiff:
				changes = append(changes, e)
				changes, b, d = diff.shiftCommons(changes, b, d)
			case PhaseInDiff:
				changes = append(changes, e)
				cc += 1
				if cc == diff.contextSize {
					phase = PhaseBehindDiff
				}
			case PhaseBehindDiff:
				// do nothing
			}
			b += 1
			d += 1
		}

		if phase == PhaseBehindDiff || i == len(diff.ses)-1 {
			a := changes[0].aIdx
			c := changes[0].bIdx
			if diff.reverse {
				a, c = c, a
			}
			switch changes[0].t {
			case SesDelete:
				a = changes[0].aIdx
				c = changes[0].aIdx
			case SesAdd:
				a = changes[0].bIdx
				c = changes[0].bIdx
			}
			uniHunk := UniHunk[T]{
				a:       a,
				b:       b,
				c:       c,
				d:       d,
				changes: changes,
			}
			uniHunks = append(uniHunks, uniHunk)
			cc = 0
			b, d = 0, 0
			changes = make([]SesElem[T], 0)
			phase = PhaseFrontDiff
		}
	}

	return uniHunks
}