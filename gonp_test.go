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
