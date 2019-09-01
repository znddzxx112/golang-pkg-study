package path

import (
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	t.Log(path.Base("/a/b/c.tmp"))
	t.Log(path.Base("../c"))

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		t.Logf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	t.Log(path.Dir("/a/b/c"))
	t.Log(path.Dir("a/b/c"))
	t.Log(path.Dir("/a/"))
	t.Log(path.Dir("a/"))
	t.Log(path.Dir("/"))
	t.Log(path.Dir(""))

	t.Log(path.Ext("/a/b/c/bar.css"))
	t.Log(path.Ext("/"))
	t.Log(path.Ext(""))

	t.Log(path.IsAbs("a/b"))
	t.Log(path.IsAbs("/a/b/c"))

	t.Log(path.Join("a/c.txt", "d"))

	t.Log(path.Match("a*", "bac/s"))
	t.Log(path.Match("a*", "bacs"))
	t.Log(path.Match("a*", "acs"))

	t.Log(path.Split("/a/b/c"))
	t.Log(path.Split("/a/b/c/"))
}
