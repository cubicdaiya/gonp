package gonp

import (
	"testing"
)

func equalsSlice[T Elem](slice1, slice2 []T) bool {
	m, n := len(slice1), len(slice2)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func equalsSesElemSlice[T Elem](ses1, ses2 []SesElem[T]) bool {
	m, n := len(ses1), len(ses2)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if ses1[i] != ses2[i] {
			return false
		}
	}
	return true
}

func equalsUniHunks[T Elem](uniHunks1, uniHunks2 []UniHunk[T]) bool {
	m, n := len(uniHunks1), len(uniHunks2)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if uniHunks1[i].a != uniHunks2[i].a {
			return false
		}
		if uniHunks1[i].b != uniHunks2[i].b {
			return false
		}
		if uniHunks1[i].c != uniHunks2[i].c {
			return false
		}
		if uniHunks1[i].d != uniHunks2[i].d {
			return false
		}
		if !equalsSesElemSlice(uniHunks1[i].changes, uniHunks2[i].changes) {
			return false
		}
	}
	return true
}

func TestStringDiff(t *testing.T) {

	tests := []struct {
		name     string
		a        string
		b        string
		ed       int
		lcs      string
		ses      []SesElem[rune]
		uniHunks []UniHunk[rune]
	}{
		{
			name: "string diff1",
			a:    "abc",
			b:    "abd",
			ed:   2,
			lcs:  "ab",
			ses: []SesElem[rune]{
				{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
				{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
				{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
				{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 3, c: 1, d: 3,
					changes: []SesElem[rune]{
						{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
						{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
						{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
						{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
					},
				},
			},
		},
		{
			name: "string diff2",
			a:    "abcdef",
			b:    "dacfea",
			ed:   6,
			lcs:  "acf",
			ses: []SesElem[rune]{
				{e: 'd', t: SesAdd, aIdx: 0, bIdx: 1},
				{e: 'a', t: SesCommon, aIdx: 1, bIdx: 2},
				{e: 'b', t: SesDelete, aIdx: 2, bIdx: 0},
				{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
				{e: 'd', t: SesDelete, aIdx: 4, bIdx: 0},
				{e: 'e', t: SesDelete, aIdx: 5, bIdx: 0},
				{e: 'f', t: SesCommon, aIdx: 6, bIdx: 4},
				{e: 'e', t: SesAdd, aIdx: 0, bIdx: 5},
				{e: 'a', t: SesAdd, aIdx: 0, bIdx: 6},
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 6, c: 1, d: 6,
					changes: []SesElem[rune]{
						{e: 'd', t: SesAdd, aIdx: 0, bIdx: 1},
						{e: 'a', t: SesCommon, aIdx: 1, bIdx: 2},
						{e: 'b', t: SesDelete, aIdx: 2, bIdx: 0},
						{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
						{e: 'd', t: SesDelete, aIdx: 4, bIdx: 0},
						{e: 'e', t: SesDelete, aIdx: 5, bIdx: 0},
						{e: 'f', t: SesCommon, aIdx: 6, bIdx: 4},
						{e: 'e', t: SesAdd, aIdx: 0, bIdx: 5},
						{e: 'a', t: SesAdd, aIdx: 0, bIdx: 6},
					},
				},
			},
		},
		{
			name: "string diff3",
			a:    "acbdeacbed",
			b:    "acebdabbabed",
			ed:   6,
			lcs:  "acbdabed",
			ses: []SesElem[rune]{
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
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 10, c: 1, d: 12,
					changes: []SesElem[rune]{
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
					},
				},
			},
		},
		{
			name: "string diff4",
			a:    "abcbda",
			b:    "bdcaba",
			ed:   4,
			lcs:  "bcba",
			ses: []SesElem[rune]{
				{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
				{e: 'b', t: SesCommon, aIdx: 2, bIdx: 1},
				{e: 'd', t: SesAdd, aIdx: 0, bIdx: 2},
				{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
				{e: 'a', t: SesAdd, aIdx: 0, bIdx: 4},
				{e: 'b', t: SesCommon, aIdx: 4, bIdx: 5},
				{e: 'd', t: SesDelete, aIdx: 5, bIdx: 0},
				{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 6, c: 1, d: 6,
					changes: []SesElem[rune]{
						{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
						{e: 'b', t: SesCommon, aIdx: 2, bIdx: 1},
						{e: 'd', t: SesAdd, aIdx: 0, bIdx: 2},
						{e: 'c', t: SesCommon, aIdx: 3, bIdx: 3},
						{e: 'a', t: SesAdd, aIdx: 0, bIdx: 4},
						{e: 'b', t: SesCommon, aIdx: 4, bIdx: 5},
						{e: 'd', t: SesDelete, aIdx: 5, bIdx: 0},
						{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
					},
				},
			},
		},
		{
			name: "string diff5",
			a:    "bokko",
			b:    "bokkko",
			ed:   1,
			lcs:  "bokko",
			ses: []SesElem[rune]{
				{e: 'b', t: SesCommon, aIdx: 1, bIdx: 1},
				{e: 'o', t: SesCommon, aIdx: 2, bIdx: 2},
				{e: 'k', t: SesCommon, aIdx: 3, bIdx: 3},
				{e: 'k', t: SesCommon, aIdx: 4, bIdx: 4},
				{e: 'k', t: SesAdd, aIdx: 0, bIdx: 5},
				{e: 'o', t: SesCommon, aIdx: 5, bIdx: 6},
			},
			uniHunks: []UniHunk[rune]{
				{a: 2, b: 4, c: 2, d: 5,
					changes: []SesElem[rune]{
						{e: 'o', t: SesCommon, aIdx: 2, bIdx: 2},
						{e: 'k', t: SesCommon, aIdx: 3, bIdx: 3},
						{e: 'k', t: SesCommon, aIdx: 4, bIdx: 4},
						{e: 'k', t: SesAdd, aIdx: 0, bIdx: 5},
						{e: 'o', t: SesCommon, aIdx: 5, bIdx: 6},
					},
				},
			},
		},
		{
			name: "string diff6",
			a:    "abcaaaaaabd",
			b:    "abdaaaaaabc",
			ed:   4,
			lcs:  "abaaaaaab",
			ses: []SesElem[rune]{
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
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 6, c: 1, d: 6,
					changes: []SesElem[rune]{
						{e: 'a', t: SesCommon, aIdx: 1, bIdx: 1},
						{e: 'b', t: SesCommon, aIdx: 2, bIdx: 2},
						{e: 'c', t: SesDelete, aIdx: 3, bIdx: 0},
						{e: 'd', t: SesAdd, aIdx: 0, bIdx: 3},
						{e: 'a', t: SesCommon, aIdx: 4, bIdx: 4},
						{e: 'a', t: SesCommon, aIdx: 5, bIdx: 5},
						{e: 'a', t: SesCommon, aIdx: 6, bIdx: 6},
					},
				},
				{a: 8, b: 4, c: 8, d: 4,
					changes: []SesElem[rune]{
						{e: 'a', t: SesCommon, aIdx: 8, bIdx: 8},
						{e: 'a', t: SesCommon, aIdx: 9, bIdx: 9},
						{e: 'b', t: SesCommon, aIdx: 10, bIdx: 10},
						{e: 'd', t: SesDelete, aIdx: 11, bIdx: 0},
						{e: 'c', t: SesAdd, aIdx: 0, bIdx: 11},
					},
				},
			},
		},
		{
			name:     "empty string diff1",
			a:        "",
			b:        "",
			ed:       0,
			lcs:      "",
			ses:      []SesElem[rune]{},
			uniHunks: []UniHunk[rune]{},
		},
		{
			name: "empty string diff2",
			a:    "a",
			b:    "",
			ed:   1,
			lcs:  "",
			ses: []SesElem[rune]{
				{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 1, c: 0, d: 0, changes: []SesElem[rune]{
					{e: 'a', t: SesDelete, aIdx: 1, bIdx: 0},
				},
				},
			},
		},
		{
			name: "empty string diff3",
			a:    "",
			b:    "b",
			ed:   1,
			lcs:  "",
			ses: []SesElem[rune]{
				{e: 'b', t: SesAdd, aIdx: 0, bIdx: 1},
			},
			uniHunks: []UniHunk[rune]{
				{a: 0, b: 0, c: 1, d: 1, changes: []SesElem[rune]{
					{e: 'b', t: SesAdd, aIdx: 0, bIdx: 1},
				},
				},
			},
		},
		{
			name: "multi byte string diff",
			a:    "久保竜彦",
			b:    "久保達彦",
			ed:   2,
			lcs:  "久保彦",
			ses: []SesElem[rune]{
				{e: '久', t: SesCommon, aIdx: 1, bIdx: 1},
				{e: '保', t: SesCommon, aIdx: 2, bIdx: 2},
				{e: '竜', t: SesDelete, aIdx: 3, bIdx: 0},
				{e: '達', t: SesAdd, aIdx: 0, bIdx: 3},
				{e: '彦', t: SesCommon, aIdx: 4, bIdx: 4},
			},
			uniHunks: []UniHunk[rune]{
				{a: 1, b: 4, c: 1, d: 4, changes: []SesElem[rune]{
					{e: '久', t: SesCommon, aIdx: 1, bIdx: 1},
					{e: '保', t: SesCommon, aIdx: 2, bIdx: 2},
					{e: '竜', t: SesDelete, aIdx: 3, bIdx: 0},
					{e: '達', t: SesAdd, aIdx: 0, bIdx: 3},
					{e: '彦', t: SesCommon, aIdx: 4, bIdx: 4},
				},
				},
			},
		},
	}

	for _, tt := range tests {
		diff := New([]rune(tt.a), []rune(tt.b))
		diff.Compose()
		ed := diff.Editdistance()
		lcs := string(diff.Lcs())
		ses := diff.Ses()
		uniHunks := diff.UnifiedHunks()
		if tt.ed != ed {
			t.Fatalf(":%s:ed: want: %d, got: %d", tt.name, tt.ed, ed)
		}
		if tt.lcs != lcs {
			t.Fatalf(":%s:lcs: want: %s, got: %s", tt.name, tt.lcs, lcs)
		}
		if !equalsSesElemSlice(tt.ses, ses) {
			t.Fatalf(":%s:ses: want: %v, got: %v", tt.name, tt.ses, ses)
		}

		if !equalsUniHunks(tt.uniHunks, uniHunks) {
			t.Fatalf(":%s:uniHunks: want: %v, got: %v", tt.name, tt.uniHunks, tt.uniHunks)
		}
	}
}

func TestSliceDiff(t *testing.T) {

	tests := []struct {
		name     string
		a        []int
		b        []int
		ed       int
		lcs      []int
		ses      []SesElem[int]
		uniHunks []UniHunk[int]
	}{
		{
			name: "int slice diff",
			a:    []int{1, 2, 3},
			b:    []int{1, 5, 3},
			ed:   2,
			lcs:  []int{1, 3},
			ses: []SesElem[int]{
				{e: 1, t: SesCommon, aIdx: 1, bIdx: 1},
				{e: 2, t: SesDelete, aIdx: 2, bIdx: 0},
				{e: 5, t: SesAdd, aIdx: 0, bIdx: 2},
				{e: 3, t: SesCommon, aIdx: 3, bIdx: 3},
			},
			uniHunks: []UniHunk[int]{
				{a: 1, b: 3, c: 1, d: 3,
					changes: []SesElem[int]{
						{e: 1, t: SesCommon, aIdx: 1, bIdx: 1},
						{e: 2, t: SesDelete, aIdx: 2, bIdx: 0},
						{e: 5, t: SesAdd, aIdx: 0, bIdx: 2},
						{e: 3, t: SesCommon, aIdx: 3, bIdx: 3},
					},
				},
			},
		},
		{
			name:     "empty slice diff",
			a:        []int{},
			b:        []int{},
			ed:       0,
			lcs:      []int{},
			ses:      []SesElem[int]{},
			uniHunks: []UniHunk[int]{},
		},
	}

	for _, tt := range tests {
		diff := New(tt.a, tt.b)
		diff.Compose()
		ed := diff.Editdistance()
		lcs := diff.Lcs()
		ses := diff.Ses()
		uniHunks := diff.UnifiedHunks()
		if tt.ed != ed {
			t.Fatalf(":%s:ed: want: %d, got: %d", tt.name, tt.ed, ed)
		}
		if !equalsSlice(tt.lcs, lcs) {
			t.Fatalf(":%s:lcs: want: %v, got: %v", tt.name, tt.lcs, lcs)
		}
		if !equalsSesElemSlice(tt.ses, ses) {
			t.Fatalf(":%s:ses: want: %v, got: %v", tt.name, tt.ses, ses)
		}

		if !equalsUniHunks(tt.uniHunks, uniHunks) {
			t.Fatalf(":%s:uniHunks: want: %v, got: %v", tt.name, tt.uniHunks, tt.uniHunks)
		}
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

	if diff.Editdistance() != 2 {
		t.Fatalf("want: 2, actual: %v", diff.Editdistance())
	}

	if lcs != "ab" {
		t.Fatalf("want: ab, actual: %v", lcs)
	}

	if !equalsSesElemSlice(sesActual, sesExpected) {
		t.Fatalf("want: %v, actual: %v", sesExpected, sesActual)
	}

	uniHunksActual := diff.UnifiedHunks()
	uniHunksExpected := []UniHunk[rune]{
		{a: 1, b: 3, c: 1, d: 3, changes: sesExpected},
	}

	if !equalsUniHunks(uniHunksActual, uniHunksExpected) {
		t.Fatalf(":uniHunks: want: %v, got: %v", uniHunksExpected, uniHunksActual)
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

	if diff.Editdistance() != 2 {
		t.Fatalf("want: 2, actual: %v", diff.Editdistance())
	}

	if lcs != "" {
		t.Fatalf("want: \"\", actual: %v", lcs)
	}

	if !equalsSesElemSlice(sesActual, sesExpected) {
		t.Fatalf("want: %v, actual: %v", sesExpected, sesActual)
	}
}

func TestDiffPluralSubsequence(t *testing.T) {
	a := []rune("abcaaaaaabd")
	b := []rune("abdaaaaaabc")
	diff := New(a, b)
	diff.SetRouteSize(2) // dividing sequences forcibly
	diff.Compose()
	if diff.Editdistance() != 4 {
		t.Fatalf("want: 4, actual: %d", diff.Editdistance())
	}
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
	if actual != expected {
		t.Fatalf("want: %v, actual: %v", expected, actual)
	}
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
	if actual != expected {
		t.Fatalf("want: %v, actual: %v", expected, actual)
	}
}

func BenchmarkStringDiffCompose(b *testing.B) {
	s1 := []rune("abc")
	s2 := []rune("abd")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diff := New(s1, s2)
		diff.Compose()
	}
}

func BenchmarkStringDiffComposeIfOnlyEd(b *testing.B) {
	s1 := []rune("abc")
	s2 := []rune("abd")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diff := New(s1, s2)
		diff.OnlyEd()
		diff.Compose()
	}
}
