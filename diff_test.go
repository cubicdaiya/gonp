package gonp

import (
	"fmt"
	"testing"
)

func equalsSesElemArray[T Elem](ses1, ses2 []SesElem[T]) bool {
	m, n := len(ses1), len(ses2)
	if m != n {
		return true
	}
	for i := 0; i < m; i++ {
		if ses1[i] != ses2[i] {
			return false
		}
	}
	return true
}

func assert(t *testing.T, b bool, msg string) {
	if !b {
		t.Error(msg)
	}
}

func TestDiff1(t *testing.T) {
	a := []rune("abc")
	b := []rune("abd")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
		{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
	}
	assert(t, diff.Editdistance() == 2, fmt.Sprintf("want: 2, actual: %d", diff.Editdistance()))
	assert(t, lcs == "ab", fmt.Sprintf("want: ab, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 3, c: 1, d: 3, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 3, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 3, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiff2(t *testing.T) {
	a := []rune("abcdef")
	b := []rune("dacfea")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'd', t: SesAdd, aIdx: 0, bIdx: 1},
		{e: 'a', t: SesCommon, aIdx: 1, bIdx: 2},
		{e: 'b', t: SesDelete, aIdx: 2, bIdx: 0},
		{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
		{e: 'd', t: SesDelete, aIdx: 4, bIdx: 0},
		{e: 'e', t: SesDelete, aIdx: 5, bIdx: 0},
		{e: 'f', t: SesCommon, aIdx: 6, bIdx: 4},
		{e: 'e', t: SesAdd, aIdx: 0, bIdx: 5},
		{e: 'a', t: SesAdd, aIdx: 0, bIdx: 6},
	}
	assert(t, diff.Editdistance() == 6, fmt.Sprintf("want: 6, actual: %d", diff.Editdistance()))
	assert(t, lcs == "acf", fmt.Sprintf("want: acf, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 6, c: 1, d: 6, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 6, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 6, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiff3(t *testing.T) {
	a := []rune("acbdeacbed")
	b := []rune("acebdabbabed")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: 'c', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: 'e', t: SesAdd, aIdx: 0, bIdx: 3},
		{e: 'b', t: SesCommon, aIdx: 3, bIdx: 4},
		{e: 'd', t: SesCommon, aIdx: 4, bIdx: 5},
		{e: 'e', t: SesDelete, aIdx: 5, bIdx: 0},
		{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
		{e: 'c', t: SesDelete, aIdx: 7, bIdx: 0},
		{e: 'b', t: SesCommon, aIdx: 8, bIdx: 7},
		{e: 'b', t: SesAdd, aIdx: 0, bIdx: 8},
		{e: 'a', t: SesAdd, aIdx: 0, bIdx: 9},
		{e: 'b', t: SesAdd, aIdx: 0, bIdx: 10},
		{e: 'e', t: SesCommon, aIdx: 9, bIdx: 11},
		{e: 'd', t: SesCommon, aIdx: 10, bIdx: 12},
	}
	assert(t, diff.Editdistance() == 6, fmt.Sprintf("want: 6, actual: %d", diff.Editdistance()))
	assert(t, lcs == "acbdabed", fmt.Sprintf("want: acbdabed, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 10, c: 1, d: 12, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want:  1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 10, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want:  1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 12, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiff4(t *testing.T) {
	a := []rune("abcbda")
	b := []rune("bdcaba")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
		{e: 'b', t: SesCommon, aIdx: 2, bIdx: 1},
		{e: 'd', t: SesAdd, aIdx: 0, bIdx: 2},
		{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
		{e: 'a', t: SesAdd, aIdx: 0, bIdx: 4},
		{e: 'b', t: SesCommon, aIdx: 4, bIdx: 5},
		{e: 'd', t: SesDelete, aIdx: 5, bIdx: 0},
		{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
	}
	assert(t, diff.Editdistance() == 4, fmt.Sprintf("want: 4, actual: %d", diff.Editdistance()))
	assert(t, lcs == "bcba", fmt.Sprintf("want: bcba, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 6, c: 1, d: 6, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 6, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 6, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiff5(t *testing.T) {
	a := []rune("bokko")
	b := []rune("bokkko")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'b', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: 'o', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: 'k', t: SesCommon, aIdx: 3, bIdx: 3},
		{e: 'k', t: SesCommon, aIdx: 4, bIdx: 4},
		{e: 'k', t: SesAdd, aIdx: 0, bIdx: 5},
		{e: 'o', t: SesCommon, aIdx: 5, bIdx: 6},
	}
	assert(t, diff.Editdistance() == 1, fmt.Sprintf("want: 1, actual: %d", diff.Editdistance()))
	assert(t, lcs == "bokko", fmt.Sprintf("want: bokko, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 2, b: 4, c: 2, d: 5, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 2, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 4, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 2, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 5, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiff6(t *testing.T) {
	a := []rune("abcaaaaaabd")
	b := []rune("abdaaaaaabc")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
		{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
		{e: 'a', t: SesCommon, aIdx: 4, bIdx: 4},
		{e: 'a', t: SesCommon, aIdx: 5, bIdx: 5},
		{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
		{e: 'a', t: SesCommon, aIdx: 7, bIdx: 7},
		{e: 'a', t: SesCommon, aIdx: 8, bIdx: 8},
		{e: 'a', t: SesCommon, aIdx: 9, bIdx: 9},
		{e: 'b', t: SesCommon, aIdx: 10, bIdx: 10},
		{e: 'd', t: SesDelete, aIdx: 11, bIdx: 0},
		{e: 'c', t: SesAdd, aIdx: 0, bIdx: 11},
	}
	assert(t, diff.Editdistance() == 4, fmt.Sprintf("want: 4, actual: %d", diff.Editdistance()))
	assert(t, lcs == "abaaaaaab", fmt.Sprintf("want: abaaaaaab, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("\nwant: %v, \nactual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 6, c: 1, d: 6, changes: []SesElem[rune]{
			{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
			{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
			{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
			{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
			{e: 'a', t: SesCommon, aIdx: 4, bIdx: 4},
			{e: 'a', t: SesCommon, aIdx: 5, bIdx: 5},
			{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
		}},
		{a: 8, b: 4, c: 8, d: 4, changes: []SesElem[rune]{
			{e: 'a', t: SesCommon, aIdx: 8, bIdx: 8},
			{e: 'a', t: SesCommon, aIdx: 9, bIdx: 9},
			{e: 'b', t: SesCommon, aIdx: 10, bIdx: 10},
			{e: 'd', t: SesDelete, aIdx: 11, bIdx: 0},
			{e: 'c', t: SesAdd, aIdx: 0, bIdx: 11},
		}},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 2, actual: %d", len(uniHunksActual)))
	wantsIdx := [2][4]int{
		{1, 6, 1, 6},
		{8, 4, 8, 4},
	}
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: %d, actual: %d", wantsIdx[i][0], uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: %d, actual: %d", wantsIdx[i][1], uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: %d, actual: %d", wantsIdx[i][2], uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: %d, actual: %d", wantsIdx[i][3], uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestPluralDiff(t *testing.T) {
	a := []rune("abc")
	b := []rune("abd")
	diff := New(a, b)
	diff.SetRouteSize(1)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
		{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
	}
	assert(t, diff.Editdistance() == 2, fmt.Sprintf("want: 2, actual: %d", diff.Editdistance()))
	assert(t, lcs == "ab", fmt.Sprintf("want: ab, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 3, c: 1, d: 3, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 3, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 3, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiffEmptyString1(t *testing.T) {
	a := []rune("")
	b := []rune("")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{}
	assert(t, diff.Editdistance() == 0, fmt.Sprintf("want: 0, actual: %d", diff.Editdistance()))
	assert(t, lcs == "", fmt.Sprintf("want: \"\", actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: [], actual: %v", uniHunksActual))
}

func TestDiffEmptyString2(t *testing.T) {
	a := []rune("a")
	b := []rune("")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
	}
	assert(t, diff.Editdistance() == 1, fmt.Sprintf("want: 1, actual: %d", diff.Editdistance()))
	assert(t, lcs == "", fmt.Sprintf("want: \"\", actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 1, c: 0, d: 0, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 0, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 0, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiffEmptyString3(t *testing.T) {
	a := []rune("")
	b := []rune("b")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: 'b', t: SesAdd, aIdx: 0, bIdx: 1},
	}
	assert(t, diff.Editdistance() == 1, fmt.Sprintf("want: 1, actual: %d", diff.Editdistance()))
	assert(t, lcs == "", fmt.Sprintf("want: \"\", actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 0, b: 0, c: 1, d: 1, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 0, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 0, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiffMultiByteString(t *testing.T) {
	a := []rune("久保竜彦")
	b := []rune("久保達彦")
	diff := New(a, b)
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{
		{e: '久', t: SesCommon, aIdx: 1, bIdx: 1},
		{e: '保', t: SesCommon, aIdx: 2, bIdx: 2},
		{e: '竜', t: SesDelete, aIdx: 3, bIdx: 0},
		{e: '達', t: SesAdd, aIdx: 0, bIdx: 3},
		{e: '彦', t: SesCommon, aIdx: 4, bIdx: 4},
	}
	assert(t, diff.Editdistance() == 2, fmt.Sprintf("want: 2, actual: %d", diff.Editdistance()))
	assert(t, lcs == "久保彦", fmt.Sprintf("want: 久保彦, actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 4, c: 1, d: 4, changes: sesExpected},
	}
	assert(t, len(uniHunksActual) == len(uniHunksExpected), fmt.Sprintf("want: 1, actual: %d", len(uniHunksActual)))
	for i := 0; i < len(uniHunksActual); i++ {
		assert(t, uniHunksActual[i].a == uniHunksExpected[i].a, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].a))
		assert(t, uniHunksActual[i].b == uniHunksExpected[i].b, fmt.Sprintf("want: 4, actual: %d", uniHunksActual[i].b))
		assert(t, uniHunksActual[i].c == uniHunksExpected[i].c, fmt.Sprintf("want: 1, actual: %d", uniHunksActual[i].c))
		assert(t, uniHunksActual[i].d == uniHunksExpected[i].d, fmt.Sprintf("want: 4, actual: %d", uniHunksActual[i].d))
		assert(t, equalsSesElemArray(uniHunksActual[i].changes, uniHunksExpected[i].changes), fmt.Sprintf("want: %v, actual: %v", uniHunksExpected, uniHunksActual))
	}
}

func TestDiffOnlyEditdistance(t *testing.T) {
	a := []rune("abc")
	b := []rune("abd")
	diff := New(a, b)
	diff.OnlyEd()
	diff.Compose()
	lcs := string(diff.Lcs())
	sesActual := diff.Ses()
	sesExpected := []SesElem[rune]{}
	assert(t, diff.Editdistance() == 2, fmt.Sprintf("want: 2, actual: %d", diff.Editdistance()))
	assert(t, lcs == "", fmt.Sprintf("want: \"\", actual: %s", lcs))
	assert(t, equalsSesElemArray(sesActual, sesExpected), fmt.Sprintf("want: %v, actual: %v", sesExpected, sesActual))
}

func TestDiffPluralSubsequence(t *testing.T) {
	a := []rune("abcaaaaaabd")
	b := []rune("abdaaaaaabc")
	diff := New(a, b)
	diff.SetRouteSize(2) // dividing sequence forcibly
	diff.Compose()
	assert(t, diff.Editdistance() == 4, fmt.Sprintf("want: 4, actual: %d", diff.Editdistance()))
}

func TestDiffSprintSes(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "1", "c"}
	diff := New(a, b)
	diff.Compose()
	actual := diff.SprintSes()
	expected := ` a
-b
+1
 c
`
	assert(t, actual == expected, fmt.Sprintf("want: %v, actual: %v", expected, actual))
}

func TestDiffSprintUniHunks(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "1", "c"}
	diff := New(a, b)
	diff.Compose()
	actual := diff.SprintUniHunks(diff.UnifiedHunks())
	expected := `@@ -1,3 +1,3 @@
 a
-b
+1
 c
`
	assert(t, actual == expected, fmt.Sprintf("want: %v, actual: %v", expected, actual))
}
