package suffixarrayt

import (
	"index/suffixarray"
	"sort"
	"testing"
)

func TestSuffixarray(t *testing.T) {
	source := []byte("hello world, hello foo")
	index := suffixarray.New(source)

	offsets := index.Lookup([]byte("hello"), -1)

	sort.Ints(offsets)

	t.Logf("%v", offsets)
}
