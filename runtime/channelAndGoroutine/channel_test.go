package channelAndGoroutine

import (
	"log"
	"sync"
	"testing"
)

// Go Channel的实现
// https://studygolang.com/articles/20714

// Goroutine浅析
// https://i6448038.github.io/2017/12/04/golang-concurrency-principle/

// https://ninokop.github.io/

// Go 内存管理
// https://ninokop.github.io/2017/12/01/go-memory-model/

// Go 垃圾回收
// https://ninokop.github.io/2017/12/07/go-gc/

// Hexo博客搭建与github托管
// https://ninokop.github.io/2017/08/25/init-hexo-blog/

// go
// https://ninokop.github.io/categories/Golang/

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)

	i, ok := <-ch
	t.Log(i, ok)

	j, ok := <-ch
	t.Log(j, ok)

	select {
	case k, ok := <-ch:
		t.Log(k, ok) // ok is false
	}
}

func TestChannelNReader(t *testing.T) {
	total := 2
	stopCh := make(chan int, total)
	ch := make(chan int, 1)

	go func(ch chan int) {
		ch <- 1
		close(ch)
	}(ch)

	for i := 1; i <= total; i++ {
		go func(stopCh chan int, ch chan int) {
			defer func() {
				stopCh <- 1
			}()
			for {
				select {
				case i, ok := <-ch:
					if !ok {
						log.Println("go exit")
						return
					}
					log.Println(i)
				}
			}
		}(stopCh, ch)
	}

	for i := 1; i <= total; i++ {
		select {
		case j := <-stopCh:
			t.Log(j)
		}
	}
}

type mychan struct {
	closech sync.Mutex
	isclose bool
	ch      chan int
}

func (m *mychan) Close() {
	m.closech.Lock()
	defer m.closech.Unlock()
	close(m.ch)
}
