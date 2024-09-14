package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main33() {

	s1 := student{"wmd", 14}
	s2 := student{"wmk", 24}

	arr := []*student{&s1, &s2}

	a, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))

	mapp := make(map[string]string, 3)
	astr := "sad1"
	bstr := "sss"
	cstr := "ddd"
	var i interface{}
	i = astr
	mapp[bstr] = astr
	mapp[cstr] = i.(string)
	fmt.Println(astr)
	fmt.Println(i)
	fmt.Println(mapp)

}
