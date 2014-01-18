package gonp

import (
	"runtime"
	"testing"
)

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
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "ab")
}

func TestDiff2(t *testing.T) {
	diff := New("abcdef", "dacfea")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acf")
}

func TestDiff3(t *testing.T) {
	diff := New("acbdeacbed", "acebdabbabed")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 6)
	assert(t, lcs == "acbdabed")
}

func TestDiff4(t *testing.T) {
	diff := New("abcbda", "bdcaba")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 4)
	assert(t, lcs == "bcba")
}

func TestDiff5(t *testing.T) {
	diff := New("bokko", "bokkko")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "bokko")
}

func TestDiff6(t *testing.T) {
	diff := New("", "")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 0)
	assert(t, lcs == "")
}

func TestDiff7(t *testing.T) {
	diff := New("a", "")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
}

func TestDiff8(t *testing.T) {
	diff := New("", "b")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 1)
	assert(t, lcs == "")
}

func TestDiff9(t *testing.T) {
	diff := New("久保達彦", "久保竜彦")
	diff.Compose()
	lcs := diff.Lcs()
	assert(t, diff.Editdistance() == 2)
	assert(t, lcs == "久保彦")
}
