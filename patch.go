package gonp

import (
	"container/list"
	"fmt"
)

// Patch applies SES between a and b to seq
func (diff *Diff[T]) Patch(seq []T) []T {
	if diff.ed == 0 {
		return seq
	}

	l := list.New()
	for _, e := range seq {
		l.PushBack(e)
	}

	le := l.Front()
	for _, e := range diff.ses {
		switch e.t {
		case SesDelete:
			lea := le.Next()
			l.Remove(le)
			le = lea
		case SesAdd:
			if le == nil {
				le = l.PushBack(e.e)
				le = le.Next()
			} else {
				l.InsertBefore(e.e, le)
			}
		case SesCommon:
			le = le.Next()
		}
	}

	r := make([]T, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		r = append(r, e.Value.(T))
	}

	return r
}

// UniPatch applies unified format difference between a and b to seq
func (diff *Diff[T]) UniPatch(seq []T, uniHunks []UniHunk[T]) ([]T, error) {
	if diff.ed == 0 {
		return seq, nil
	}
	if len(uniHunks) == 0 {
		return []T{}, fmt.Errorf("invalid difference")
	}

	l := list.New()
	for _, e := range seq {
		l.PushBack(e)
	}

	le := l.Front()
	for i := 0; i < uniHunks[0].a-1; i++ {
		if le == nil {
			return []T{}, fmt.Errorf("invalid difference")
		}
		le = le.Next()
	}

	p := 0
	for _, h := range uniHunks {
		if p != 0 {
			for i := 0; i < h.a-p-1; i++ {
				if le == nil {
					return []T{}, fmt.Errorf("invalid difference")
				}
				le = le.Next()
			}
		}

		for _, e := range h.changes {
			switch e.t {
			case SesDelete:
				lea := le.Next()
				l.Remove(le)
				le = lea
			case SesAdd:
				if le == nil {
					le = l.PushBack(e.e)
					le = le.Next()
				} else {
					l.InsertBefore(e.e, le)
				}
			case SesCommon:
				le = le.Next()
			}
		}

		p = h.a + h.b - 1
	}

	r := make([]T, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		r = append(r, e.Value.(T))
	}

	return r, nil
}
