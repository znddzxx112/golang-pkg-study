package encodingt

import (
	"encoding/base64"
	"testing"
)

func TestEncode(t *testing.T) {
	dstStr := base64.StdEncoding.EncodeToString([]byte("中文aaa"))
	t.Log(dstStr)
	// Output :5Lit5paHYWFh
}

func TestDecode(t *testing.T) {
	srcStr, _ := base64.StdEncoding.DecodeString("5Lit5paHYWFh")
	t.Log(string(srcStr))
}
