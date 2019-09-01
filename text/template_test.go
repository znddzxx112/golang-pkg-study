package textt

import (
	"html/template"
	"os"
	"testing"
)

func TestJsEscaper(t *testing.T) {
	jsString := `
	<script>alert("xss")</script>
`
	t.Log(template.JSEscapeString(jsString))
}

func TestTemplate(t *testing.T) {
	text1 := `
{{.Count}} items are made of {{.Material}}
`
	tmpl, ParseErr := template.New("text1").Parse(text1)
	if ParseErr != nil {
		t.Fatal(ParseErr)
	}

	tmpl.Execute(os.Stdout, map[string]string{"Count": "foo", "Material": "bar"})
	tmpl.Execute(os.Stdout, map[string]string{"Count": "foo1", "Material": "bar2"})
}
