package gonp

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func buildLongString(prefix, interfix, suffix string, n int) string {
	var sb strings.Builder
	sb.WriteString(prefix)
	for i := 0; i < n; i++ {
		sb.WriteString(interfix)
	}
	sb.WriteString(suffix)
	return sb.String()
}

func TestPatch(t *testing.T) {

	tests := []struct {
		a string
		b string
	}{
		{
			a: "",
			b: "",
		},
		{
			a: "a",
			b: "a",
		},
		{
			a: "abc",
			b: "",
		},
		{
			a: "",
			b: "def",
		},
		{
			a: "abc",
			b: "abd",
		},
		{
			a: "abcdef",
			b: "dacfea",
		},
		{
			a: "acbdeacbed",
			b: "acebdabbabed",
		},
		{
			a: "bokko",
			b: "bokkko",
		},
		{
			a: "abcbda",
			b: "bdcaba",
		},
		{
			a: "abcaaaaaaaaaaaaaabd",
			b: "abdaaaaaaaaaaaaaabc",
		},
		{
			a: "aaaaaaaaaaaaaaaaadsafabcaaaaaaaaaaaaaaaaaaewaaabdaaaaaabbb",
			b: "aaaaaaaaaaaaaaadasfdsafsadasdafbdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaesaaabcaaaaaaccc",
		},
		{
			a: "aaaaaaaaaaaaaaaaadsafabcaaaaaaaaaaaaaaaaaaewaaabdaaaaaabbb",
			b: "aaaaaaaaaaaaaaadasfdsafsadasdafbdaaaaaaaaaaaaaaaaaaaaaaeaaaaaaaaaaesaaabcaaaaaaccc",
		},
		{
			a: "aaaaaaaaaaaaaaadasfdsafsadasdafbdaaaaaaaaaaaaaaaaaaaaaaeaaaaaaaaaaesaaabcaaaaaaccc",
			b: "aaaaaaaaaaaaaaaaadsafabcaaaaaaaaaaaaaaaaaaewaaabdaaaaaabbb",
		},
		{
			a: "aaaaaaaaaaaaa>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>aaaadsafabcaaaaaaaaaaaaaaaaaaewaaabdaaaaaabbb",
			b: "aaaaaaaaaaaaaaadasfdsafsadasdafbaaaaaaaaaaaaaaaaaeaaaaaaaaaae&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&saaabcaaaaaaccc",
		},
	}

	for _, test := range tests {
		diff := New([]rune(test.a), []rune(test.b))
		diff.Compose()

		patchedSeq := diff.Patch([]rune(test.a))
		if string(patchedSeq) != test.b {
			t.Errorf("applying SES between '%s' and '%s' to '%s' is %s, but got %s", string(test.a), string(test.b), string(test.a), string(test.b), string(patchedSeq))
		}

		uniPatchedSeq, _ := diff.UniPatch([]rune(test.a), diff.UnifiedHunks())
		if string(uniPatchedSeq) != test.b {
			t.Errorf("applying unified format difference between '%s' and '%s' to '%s' is %s, but got %s", string(test.a), string(test.b), string(test.a), string(test.b), string(uniPatchedSeq))
		}
	}
}

func BenchmarkPatchSmall(b *testing.B) {
	s1 := []rune("abc")
	s2 := []rune("abd")
	diff := New(s1, s2)
	diff.Compose()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = diff.Patch(s1)
	}
}

func BenchmarkUniPatchSmall(b *testing.B) {
	s1 := []rune("abc")
	s2 := []rune("abd")
	diff := New(s1, s2)
	diff.Compose()
	uniHunks := diff.UnifiedHunks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = diff.UniPatch(s1, uniHunks)
	}
}

func BenchmarkPatchLarge(b *testing.B) {
	s1 := []rune(buildLongString("abc", "a", "def", 1000))
	s2 := []rune(buildLongString("abd", "a", "ghi", 1000))
	diff := New([]rune(s1), []rune(s2))
	diff.Compose()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = diff.Patch(s1)
	}
}

func BenchmarkUniPatchLarge(b *testing.B) {
	s1 := []rune(buildLongString("abc", "a", "def", 1000))
	s2 := []rune(buildLongString("abd", "a", "ghi", 1000))
	diff := New(s1, s2)
	diff.Compose()
	uniHunks := diff.UnifiedHunks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = diff.UniPatch(s1, uniHunks)
	}
}

func FuzzPatch(f *testing.F) {
	f.Add("aaaaaaaaaaaaaaaaadsafabcaaaaaaaaaaaaaaaaaaewaaabdaaaaaabbb", "aaaaaaaaaaaaaaadasfdsafsadasdafbdaaaaaaaaaaaaaaaaaaaaaaeaaaaaaaaaaesaaabcaaaaaaccc")
	f.Fuzz(func(t *testing.T, a, b string) {
		if !utf8.ValidString(a) || !utf8.ValidString(b) {
			return
		}

		diff := New([]rune(a), []rune(b))
		diff.Compose()

		patchedSeq := diff.Patch([]rune(a))
		if string(patchedSeq) != b {
			t.Errorf("applying SES between '%s' and '%s' to '%s' is %s, but got %s", a, b, a, b, string(patchedSeq))
		}

		uniPatchedSeq, _ := diff.UniPatch([]rune(a), diff.UnifiedHunks())
		if string(uniPatchedSeq) != b {
			t.Errorf("applying unified format difference between '%s' and '%s' to '%s' is %s, but got %s", a, b, a, b, string(uniPatchedSeq))
		}
	})
}
