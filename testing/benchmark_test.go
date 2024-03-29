package testingt

import (
	"bytes"
	"html/template"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 1; i < b.N; i++ {
		//fmt.Println(i)
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	b.ReportAllocs()
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		// 每个 goroutine 有属于自己的 bytes.Buffer.
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
