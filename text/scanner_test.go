package textt

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/scanner"
	"text/tabwriter"
)

var longtext = `
NewReaderSize 将 rd 封装成一个带缓存的 bufio.Reader 对象，
// 缓存大小由 size 指定（如果小于 16 则会被设置为 16）。
// 如果 rd 的基类型就是有足够缓存的 bufio.Reader 类型，则直接将
// rd 转换为基类型返回。
func NewReaderSize
`

func TestScanner(t *testing.T) {
	var scan scanner.Scanner
	bsReader := strings.NewReader(longtext)
	sc := scan.Init(bsReader)

	var tok rune
	for tok = sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		if tok == '\n' {
			t.Log("whitespace")
		}
		t.Logf("%s:%s", sc.Pos(), sc.TokenText())
	}

}

func TestTabWriter(t *testing.T) {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}
