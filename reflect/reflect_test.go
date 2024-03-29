package reflectt

import (
	"reflect"
	"testing"
)

type foo struct {
	id   string
	name string
}

func (f *foo) String() string {
	return f.id + f.name
}
func TestTypeOf(t *testing.T) {

	for _, v := range []interface{}{
		8, int32(32), int64(64),
		"hello reflect", true, foo{id: "1", name: "lee"}} {
		oftype := reflect.TypeOf(v)
		t.Log(oftype.Kind().String())
		t.Log(oftype.Size())
		t.Log(oftype.Name())
		t.Log(oftype.NumMethod())
	}

}

func TestValueOf(t *testing.T) {
	for _, v := range []interface{}{
		8, int32(32), int64(64),
		"hello reflect", true, foo{id: "1", name: "lee"}} {
		oftype := reflect.ValueOf(v)
		t.Log(oftype.Kind().String())
		t.Log(oftype.String())
	}
}
