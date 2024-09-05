package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"time"
)

func main24() { //并发map

	m := cmap.New[int]()
	m.Set("124", 124)
	if v, ok := m.Get("124"); ok {
		fmt.Println(v)
	}
	a := func(str string) {
		for i := 0; i < 10; i++ {
			m.Set(fmt.Sprintf("%s%d", str, i), i)
		}
	}
	go a("mk1")
	go a("mk2")

	time.Sleep(time.Second * 2)
	fmt.Println(m.Count())
	fmt.Println(m.Items())
}
