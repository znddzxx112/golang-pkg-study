package containert

import (
	"container/ring"
	"testing"
)

func TestRing(t *testing.T) {
	r1 := ring.New(5)
	t.Log(r1.Len())

	r1.Value = 1
	r2 := r1.Next()
	r2.Value = 2
	r3 := r2.Next()
	r3.Value = 3
	r4 := r3.Next()
	r4.Value = 4
	r5 := r4.Next()
	r5.Value = 5

	r3.Do(func(i interface{}) {
		t.Log(i.(int))
	})

	rm := r1.Move(7)
	t.Log(rm.Value)

	t.Log("========")

	lr := r1.Link(r3)
	lr.Do(func(i interface{}) {
		t.Log(i)
	})

	t.Log("========")
	r1.Do(func(i interface{}) {
		t.Log(i)
	})

	nr := ring.New(2)
	nr.Value = 6
	nr1 := nr.Next()
	nr1.Value = 7

	r1 = r1.Link(nr)
	t.Log("========")
	r1.Do(func(i interface{}) {
		t.Log(i)
	})

	r1.Unlink(2)
	t.Log("========")
	r1.Do(func(i interface{}) {
		t.Log(i)
	})
}
