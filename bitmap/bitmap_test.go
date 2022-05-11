package bitmap

import (
	"testing"
)

func TestNew(t *testing.T) {
	bm := New(1024)
	bm.Set(10)
	t.Log(bm.Peek(10) == true)
}
