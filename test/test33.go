package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func main() {

	s1 := student{"wmd", 14}
	s2 := student{"wmk", 24}

	arr := []*student{&s1, &s2}

	a, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))
}
