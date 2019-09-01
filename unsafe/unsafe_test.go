package unsafet

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	var i1 int64 = 1
	var i2 int8 = 2

	t.Log(unsafe.Sizeof(i1))
	t.Log(unsafe.Sizeof(i2))
}
