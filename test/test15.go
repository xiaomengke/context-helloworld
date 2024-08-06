package main

import "fmt"

//// 错误示例
//func main() {
//	var data interface{} = "great"
//
//	// data 混用
//	if data, ok := data.(int); ok {
//		fmt.Println("[is an int], data: ", data)
//	} else {
//		fmt.Println("[not an int], data: ", data) // [isn't a int], data:  0
//	}
//}

// 正确示例
func main15() {
	var data interface{} = "great"

	if res, ok := data.(int); ok {
		fmt.Println("[is an int], data: ", res)
	} else {
		fmt.Println("[not an int], data: ", data) // [not an int], data:  great
	}
}
