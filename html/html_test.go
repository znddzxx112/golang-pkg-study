package htmlt

import (
	"html"
	"testing"
)

func TestHtml(t *testing.T) {
	s := `<script>alter("xss")</script>`
	es := html.EscapeString(s)
	t.Log(es)
	t.Log(html.UnescapeString(es))
}
