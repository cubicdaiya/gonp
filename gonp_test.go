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
		if ses1[i].e != ses2[i].e || ses1[i].t != ses2[i].t {
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
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'a', t: SesCommon},
		{e: 'b', t: SesCommon},
		{e: 'c', t: SesDelete},
		{e: 'd', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "ab")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff2(t *testing.T) {
	diff := New("abcdef", "dacfea")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'd', t: SesAdd},
		{e: 'a', t: SesCommon},
		{e: 'b', t: SesDelete},
		{e: 'c', t: SesCommon},
		{e: 'd', t: SesDelete},
		{e: 'e', t: SesDelete},
		{e: 'f', t: SesCommon},
		{e: 'e', t: SesAdd},
		{e: 'a', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acf")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff3(t *testing.T) {
	diff := New("acbdeacbed", "acebdabbabed")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'a', t: SesCommon},
		{e: 'c', t: SesCommon},
		{e: 'e', t: SesAdd},
		{e: 'b', t: SesCommon},
		{e: 'd', t: SesCommon},
		{e: 'e', t: SesDelete},
		{e: 'a', t: SesCommon},
		{e: 'c', t: SesDelete},
		{e: 'b', t: SesCommon},
		{e: 'b', t: SesAdd},
		{e: 'a', t: SesAdd},
		{e: 'b', t: SesAdd},
		{e: 'e', t: SesCommon},
		{e: 'd', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acbdabed")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff4(t *testing.T) {
	diff := New("abcbda", "bdcaba")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'a', t: SesDelete},
		{e: 'b', t: SesCommon},
		{e: 'd', t: SesAdd},
		{e: 'c', t: SesCommon},
		{e: 'a', t: SesAdd},
		{e: 'b', t: SesCommon},
		{e: 'd', t: SesDelete},
		{e: 'a', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 4)
	assert(t, lcs == "bcba")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiff5(t *testing.T) {
	diff := New("bokko", "bokkko")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'b', t: SesCommon},
		{e: 'o', t: SesCommon},
		{e: 'k', t: SesCommon},
		{e: 'k', t: SesCommon},
		{e: 'k', t: SesAdd},
		{e: 'o', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "bokko")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString1(t *testing.T) {
	diff := New("", "")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{}
	assert(t, diff.Editdistance() == 0)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString2(t *testing.T) {
	diff := New("a", "")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'a', t: SesDelete},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffEmptyString3(t *testing.T) {
	diff := New("", "b")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: 'b', t: SesAdd},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffMultiByteString(t *testing.T) {
	diff := New("久保竜彦", "久保達彦")
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{
		{e: '久', t: SesCommon},
		{e: '保', t: SesCommon},
		{e: '竜', t: SesDelete},
		{e: '達', t: SesAdd},
		{e: '彦', t: SesCommon},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "久保彦")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}

func TestDiffOnlyEditdistance(t *testing.T) {
	diff := New("abc", "abd")
	diff.OnlyEd()
	diff.Compose()
	lcs := diff.LcsString()
	sesActual := diff.Ses()
	sesExpected := []SesElem{}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(sesActual, sesExpected))
}
