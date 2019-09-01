package runtimet

import (
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	t.Log(runtime.Caller(0))
	t.Log(runtime.Caller(1))
	t.Log(runtime.Caller(2))
	t.Log(runtime.Caller(3))
	t.Log(runtime.Caller(4))
}
