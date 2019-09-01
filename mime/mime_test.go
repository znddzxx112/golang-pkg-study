package mimet

import (
	"mime"
	"testing"
)

func TestMime(t *testing.T) {
	t.Log(mime.ParseMediaType("image/png"))
	t.Log(mime.FormatMediaType("image/gif", nil))

	t.Log(mime.TypeByExtension(".png"))
	t.Log(mime.ExtensionsByType("image/png"))
}
