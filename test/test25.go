package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync"
	"time"
)

var s sync.RWMutex
var w sync.WaitGroup

func main25() {
	mapTest()
	syncMapTest()
	concurMapTest()
}
func mapTest() {
	m := map[string]int{"1": 1}
	startTime := time.Now().Nanosecond()
	w.Add(1)
	go writeMap(m)
	w.Add(1)
	go writeMap(m)
	//w.Add(1)
	//go readMap(m)

	w.Wait()
	endTime := time.Now().Nanosecond()
	timeDiff := endTime - startTime
	fmt.Println("map:", timeDiff)
}

func writeMap(m map[string]int) {
	defer w.Done()
	i := 0
	for i < 100000 {
		// 加锁
		s.Lock()
		m[fmt.Sprintf("%d", i)] = 1
		// 解锁
		s.Unlock()
		i++
	}
}

func readMap(m map[int]int) {
	defer w.Done()
	i := 0
	for i < 10000 {
		s.RLock()
		_ = m[1]
		s.RUnlock()
		i++
	}
}

func syncMapTest() {
	m := sync.Map{}
	//m.Store(fmt.Sprintf("%d", 1), 1)
	startTime := time.Now().Nanosecond()
	w.Add(1)
	go writeSyncMap(m)
	w.Add(1)
	go writeSyncMap(m)
	//w.Add(1)
	//go readSyncMap(m)

	w.Wait()
	endTime := time.Now().Nanosecond()
	timeDiff := endTime - startTime
	fmt.Println("sync.Map:", timeDiff)
}

func writeSyncMap(m sync.Map) {
	defer w.Done()
	i := 0
	for i < 100000 {
		m.Store(fmt.Sprintf("%d", i), 1)
		i++
	}
}

func readSyncMap(m sync.Map) {
	defer w.Done()
	i := 0
	for i < 10000 {
		m.Load(1)
		i++
	}
}

func concurMapTest() {
	m := cmap.New[int]()
	//m.Set(1, 1)
	startTime := time.Now().Nanosecond()
	w.Add(1)
	go writeConcurMap(m)
	w.Add(1)
	go writeConcurMap(m)
	w.Wait()
	endTime := time.Now().Nanosecond()
	timeDiff := endTime - startTime
	fmt.Println("concurMap:", timeDiff)

}
func writeConcurMap(m cmap.ConcurrentMap[string, int]) {
	defer w.Done()
	i := 0
	for i < 100000 {
		m.Set(fmt.Sprintf("%d", i), 1)
		i++
	}

}
