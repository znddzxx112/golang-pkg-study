package testingt

import (
	"fmt"
	"testing"
)

// this test is ignore
func TestLogAndSkip(t *testing.T) {
	t.Log(t.Name())
	t.Log("print log")
	t.Skip("is skip")
	t.Logf("%s", "not print")
	t.SkipNow()
	t.Logf("%s", " NOT printf log")
}

// this test is fail but continue to exec
func TestError(t *testing.T) {
	t.Error("print error") // log + fail
	t.Log("is exec")
}

// this test is fail and not continue to exec
func TestFatal(t *testing.T) {
	t.Fatal("print fatal")
	t.Log("not exec")
}

func calSum(a, b int) int {
	return a + b
}

// "Output:" and is compared with
// the standard output of the function
func ExampleCalSum() {
	fmt.Print(calSum(2, 3))
	// Output:
	// 5
}
