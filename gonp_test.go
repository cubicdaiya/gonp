package gonp

import (
	"testing"
)

func equalsSesElemArray(ses1, ses2 []SesElem) bool {
	m, n := len(ses1), len(ses2)
	if m != n {
		return true
	}
	for i := 0; i < m; i++ {
		if ses1[i].c != ses2[i].c || ses1[i].t != ses2[i].t {
			return false
		}
	}
	return true
}

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestDiff1(t *testing.T) {
	diff := New("abc", "abd")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'a', t: SesCommon},
		{c: 'b', t: SesCommon},
		{c: 'c', t: SesDelete},
		{c: 'd', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "ab")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff2(t *testing.T) {
	diff := New("abcdef", "dacfea")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'd', t: SesAdd},
		{c: 'a', t: SesCommon},
		{c: 'b', t: SesDelete},
		{c: 'c', t: SesCommon},
		{c: 'd', t: SesDelete},
		{c: 'e', t: SesDelete},
		{c: 'f', t: SesCommon},
		{c: 'e', t: SesAdd},
		{c: 'a', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acf")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff3(t *testing.T) {
	diff := New("acbdeacbed", "acebdabbabed")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'a', t: SesCommon},
		{c: 'c', t: SesCommon},
		{c: 'e', t: SesAdd},
		{c: 'b', t: SesCommon},
		{c: 'd', t: SesCommon},
		{c: 'e', t: SesDelete},
		{c: 'a', t: SesCommon},
		{c: 'c', t: SesDelete},
		{c: 'b', t: SesCommon},
		{c: 'b', t: SesAdd},
		{c: 'a', t: SesAdd},
		{c: 'b', t: SesAdd},
		{c: 'e', t: SesCommon},
		{c: 'd', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acbdabed")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff4(t *testing.T) {
	diff := New("abcbda", "bdcaba")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'a', t: SesDelete},
		{c: 'b', t: SesCommon},
		{c: 'd', t: SesAdd},
		{c: 'c', t: SesCommon},
		{c: 'a', t: SesAdd},
		{c: 'b', t: SesCommon},
		{c: 'd', t: SesDelete},
		{c: 'a', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 4)
	assert(t, lcs == "bcba")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff5(t *testing.T) {
	diff := New("bokko", "bokkko")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'b', t: SesCommon},
		{c: 'o', t: SesCommon},
		{c: 'k', t: SesCommon},
		{c: 'k', t: SesCommon},
		{c: 'k', t: SesAdd},
		{c: 'o', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "bokko")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString1(t *testing.T) {
	diff := New("", "")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{}
	assert(t, diff.Editdistance() == 0)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString2(t *testing.T) {
	diff := New("a", "")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'a', t: SesDelete},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString3(t *testing.T) {
	diff := New("", "b")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: 'b', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffMultiByteString(t *testing.T) {
	diff := New("久保竜彦", "久保達彦")
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{c: '久', t: SesCommon},
		{c: '保', t: SesCommon},
		{c: '竜', t: SesDelete},
		{c: '達', t: SesAdd},
		{c: '彦', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "久保彦")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffOnlyEditdistance(t *testing.T) {
	diff := New("abc", "abd")
	diff.OnlyEd()
	diff.Compose()
	lcs := diff.Lcs()
	sesActual := diff.Ses()
	sesExpected := []SesElem{}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}
