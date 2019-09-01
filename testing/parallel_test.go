package testingt

import (
	"fmt"
	"testing"
)

var chanInt chan int = make(chan int)

func TestParall1(t *testing.T) {
	t.Parallel()
	fmt.Println("p1")
	chanInt <- 1
	fmt.Println("p1")
	chanInt <- 2
}

func TestParall2(t *testing.T) {
	t.Parallel()
	fmt.Println(<-chanInt)
	fmt.Println("p2")
	fmt.Println(<-chanInt)
	fmt.Println("p2")
}

func TestSubtests(t *testing.T) {
	t.Run("", TestParall1)
	t.Run("", TestParall2)
}
