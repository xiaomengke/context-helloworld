package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type Human struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Profile  string `structs:"profile"`
	IsGopher bool   `json:"isGopher"`
}

func main10() {
	human := Human{"Detector", 18, "A tester", true} //struct转map
	fmt.Println("========")
	fmt.Println("Third lb：", structs.Map(human))
}
