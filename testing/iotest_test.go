package testingt

import (
	"bytes"
	"testing"
	"testing/iotest"
)

func TestDataErrReader(t *testing.T) {
	bs := []byte("sasdfasdf")
	bsReader := bytes.NewReader(bs)
	bsErrReader := iotest.DataErrReader(bsReader)

	fbs := make([]byte, 50)
	fbsINt, ReadAllErr := bsErrReader.Read(fbs)
	if ReadAllErr == nil {
		t.Log(string(fbs[:fbsINt]))
		t.Fatal(ReadAllErr)
	}
	t.Log(ReadAllErr)

}

func TestHalfReader(t *testing.T) {
	bs := []byte("sasdfasdf")
	bsReader := bytes.NewReader(bs)

	halfReader := iotest.HalfReader(bsReader)
	fbs := make([]byte, 50)
	fbsNum, ReadErr := halfReader.Read(fbs)
	if ReadErr != nil {
		t.Fatal(ReadErr)
	}
	t.Log(string(fbs[:fbsNum]))
}
