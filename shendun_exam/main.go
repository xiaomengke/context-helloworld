package main

import (
	"fmt"
)

func main() {
	GenerateOrder()
	var id int
	for {
		fmt.Println("tips:input 0 to exit")
		fmt.Println("input id to query, id:")
		_, err := fmt.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("id:", id)
		if id == 0 {
			break
		}
		rt, err := QueryOrderById(id)
		if err != nil {
			fmt.Println("111")
			fmt.Println(err)
		}
		fmt.Println(rt)
	}
}
