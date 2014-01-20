package gonp

import (
	"runtime"
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

func init() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
}

func TestDiff1(t *testing.T) {
	diff := New("abc", "abd")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'a', t: Common},
		{c: 'b', t: Common},
		{c: 'c', t: Delete},
		{c: 'd', t: Add},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "ab")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiff2(t *testing.T) {
	diff := New("abcdef", "dacfea")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'd', t: Add},
		{c: 'a', t: Common},
		{c: 'b', t: Delete},
		{c: 'c', t: Common},
		{c: 'd', t: Delete},
		{c: 'e', t: Delete},
		{c: 'f', t: Common},
		{c: 'e', t: Add},
		{c: 'a', t: Add},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acf")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiff3(t *testing.T) {
	diff := New("acbdeacbed", "acebdabbabed")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'a', t: Common},
		{c: 'c', t: Common},
		{c: 'e', t: Add},
		{c: 'b', t: Common},
		{c: 'd', t: Common},
		{c: 'e', t: Delete},
		{c: 'a', t: Common},
		{c: 'c', t: Delete},
		{c: 'b', t: Common},
		{c: 'b', t: Add},
		{c: 'a', t: Add},
		{c: 'b', t: Add},
		{c: 'e', t: Common},
		{c: 'd', t: Common},
	}
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acbdabed")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiff4(t *testing.T) {
	diff := New("abcbda", "bdcaba")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'a', t: Delete},
		{c: 'b', t: Common},
		{c: 'd', t: Add},
		{c: 'c', t: Common},
		{c: 'a', t: Add},
		{c: 'b', t: Common},
		{c: 'd', t: Delete},
		{c: 'a', t: Common},
	}
	assert(t, diff.Editdistance() == 4)
	assert(t, lcs == "bcba")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiff5(t *testing.T) {
	diff := New("bokko", "bokkko")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'b', t: Common},
		{c: 'o', t: Common},
		{c: 'k', t: Common},
		{c: 'k', t: Common},
		{c: 'k', t: Add},
		{c: 'o', t: Common},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "bokko")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiffEmptyString1(t *testing.T) {
	diff := New("", "")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{}
	assert(t, diff.Editdistance() == 0)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiffEmptyString2(t *testing.T) {
	diff := New("a", "")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'a', t: Delete},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiffEmptyString3(t *testing.T) {
	diff := New("", "b")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: 'b', t: Add},
	}
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiffMultiByteString(t *testing.T) {
	diff := New("久保竜彦", "久保達彦")
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{
		{c: '久', t: Common},
		{c: '保', t: Common},
		{c: '竜', t: Delete},
		{c: '達', t: Add},
		{c: '彦', t: Common},
	}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "久保彦")
	assert(t, equalsSesElemArray(ses, ses_))
}

func TestDiffOnlyEditdistance(t *testing.T) {
	diff := New("abc", "abd")
	diff.OnlyEd()
	diff.Compose()
	lcs := diff.Lcs()
	ses := diff.Ses()
	ses_ := []SesElem{}
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "")
	assert(t, equalsSesElemArray(ses, ses_))
}
