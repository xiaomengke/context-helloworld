package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func main36() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	chanLeakOfMemory()
	time.Sleep(time.Second * 3) // 等待 goroutine 执行，防止过早输出结果
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

}
func chanLeakOfMemory() {
	errCh := make(chan error)
	go func() {
		time.Sleep(2 * time.Second)
		errCh <- errors.New("chan error") // (1)
		fmt.Println("finish sending")
	}()

	var err error
	select {
	case <-time.After(time.Second * 3): // (2) 大家也经常在这里使用 <-ctx.Done()
		fmt.Println("超时")
	case err = <-errCh:
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(nil)
		}
	}
}
