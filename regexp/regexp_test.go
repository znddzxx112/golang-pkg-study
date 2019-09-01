package regexpt

import (
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	t.Log(regexp.MatchString(`1\d2`, "zzz123zzz"))
}

func TestQuoteMeta(t *testing.T) {
	t.Log(regexp.QuoteMeta(`1\d2`))
}

// https://www.cnblogs.com/yalibuxiao/p/4194881.html
// posix和perl标准的正则表达式区别;
func TestCompile(t *testing.T) {
	pat := `\d+`
	reg, CompileErr := regexp.Compile(pat)
	if CompileErr != nil {
		t.Fatal(CompileErr)
	}

	t.Log(reg.String())

	bs := "zzz1234zz11z"
	t.Log(reg.FindString(bs))

	reg.Longest()
	t.Log(reg.MatchString(bs))
	t.Log(reg.FindAllString(bs, -1))
	t.Log(reg.ReplaceAllString(bs, "bbbbb"))

}
