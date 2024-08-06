package main

import (
	"fmt"
	"time"
)

func main5() {
	ch := make(chan string)

	go func() {
		//time.Sleep(1 * time.Second) // 模拟需要长时间运行的操作
		for m := range ch {
			fmt.Println("Processed:", m)
			time.Sleep(1 * time.Second) // 模拟需要长时间运行的操作
		}
	}()
	//fmt.Println("111")
	ch <- "cmd.1"
	//fmt.Println("222")
	ch <- "cmd.2" // 不会被接收处理
	//time.Sleep(1 * time.Second) // 模拟需要长时间运行的操作

}
