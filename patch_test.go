package gonp

import (
	"testing"
)

func TestPatch(t *testing.T) {

	tests := []struct {
		a string
		b string
	}{
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
	}

	for _, test := range tests {
		diff := New([]rune(test.a), []rune(test.b))
		diff.Compose()

		patchedSeq, _ := diff.Patch([]rune(test.a))
		if string(patchedSeq) != test.b {
			t.Errorf("applying SES between '%s' and '%s' to '%s' is %s, but got %s", string(test.a), string(test.b), string(test.a), string(test.b), string(patchedSeq))
		}

		uniPatchedSeq, _ := diff.UniPatch([]rune(test.a), diff.UnifiedHunks())
		if string(uniPatchedSeq) != test.b {
			t.Errorf("applying unified format difference between '%s' and '%s' to '%s' is %s, but got %s", string(test.a), string(test.b), string(test.a), string(test.b), string(uniPatchedSeq))
		}
	}
}
