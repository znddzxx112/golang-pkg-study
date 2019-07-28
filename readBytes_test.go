package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var Input string

func TestReadBytes(t *testing.T) {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		fmt.Print("请输入一些字符串>")
		Input, _ = f.ReadString('\n')
		if len(Input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		fmt.Printf("您输入的是:%s",Input)
	}
}
