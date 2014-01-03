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

func TestEditdistance(t *testing.T) {
	diff := New("abc", "abd")
	diff.Compose()
	assert(t, diff.Editdistance() == 2)
}
