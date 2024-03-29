package contextt

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func genInt(ctx context.Context, chanInt chan<- int) {
	var n int = 0
	var isStart bool = true
	for isStart {
		select {
		case <-ctx.Done():
			close(chanInt)
			isStart = false
		default:
			chanInt <- n
		}
		n++
	}
	return
}

func TestCtxCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	chanInt := make(chan int, 1)
	go genInt(ctx, chanInt)

	var isStart bool = true
	for isStart {
		select {
		case i, ok := <-chanInt:
			if !ok {
				isStart = false
			}
			fmt.Println(i)
			if i == 5 {
				cancel()
			}
		}
	}

}

func toRequest(ctx context.Context, chanInt chan<- int) {
	var n int = 11
	var isStart bool = true
	for isStart {
		select {
		case <-ctx.Done():
			DeadlineTime, ok := ctx.Deadline()
			fmt.Println(DeadlineTime.Format(time.RFC3339), ok)
			fmt.Println(ctx.Err())
			chanInt <- n
			isStart = false
		}
	}
}

func toRequest2(ctx context.Context, chanInt chan<- int) {
	var n int = 12
	var isStart bool = true
	for isStart {
		select {
		case <-ctx.Done():
			DeadlineTime, ok := ctx.Deadline()
			fmt.Println(DeadlineTime.Format(time.RFC3339), ok)
			fmt.Println(ctx.Err())
			chanInt <- n
			isStart = false
		}
	}
}

func TestCtxCancelMutli(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	chanInt := make(chan int, 1)
	go toRequest(ctx, chanInt)

	chanInt2 := make(chan int, 1)
	go toRequest2(ctx, chanInt2)

	cancel()
	fmt.Println(<-chanInt)
	fmt.Println(<-chanInt2)

}

func TestCtxTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	chanInt := make(chan int, 1)
	go toRequest(ctx, chanInt)

	chanInt2 := make(chan int, 1)
	go toRequest2(ctx, chanInt2)

	fmt.Println(<-chanInt)
	fmt.Println(<-chanInt2)
}
