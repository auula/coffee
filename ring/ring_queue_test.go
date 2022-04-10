package ring

import "testing"

func TestRing(t *testing.T) {
	r := New[int](20)

	for i := 0; i < 10; i++ {
		r.Put(i)
	}

	t.Log("Size:", r.Size())
	t.Log("IsFull:", r.IsFull())

	for i := 0; i < 10; i++ {
		t.Log(*r.Get())
	}

	t.Log("Size:", r.Size())
	t.Log("IsFull:", r.IsFull())

}
