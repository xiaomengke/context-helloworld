package main

import "fmt"

type TransInfo struct {
	int
}

type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

func main26() {
	var f1 Fragment = new(GetPodAction)
	var f2 Fragment = GetPodAction{}
	var f3 Fragment = &GetPodAction{}
	_ = f1
	_ = f2
	_ = f3

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)

}
