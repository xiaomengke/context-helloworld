package main

import "fmt"

//	func main() {
//		var data *byte
//		var in interface{}
//		fmt.Printf("%p\n", data)
//		fmt.Printf("%p\n", in)
//		fmt.Println(data, data == nil) // <nil> true
//		fmt.Println(in, in == nil)     // <nil> true
//
//		in = data
//		fmt.Println(in, in == nil) // <nil> false    // data 值为 nil，但 in 值不为 nil
//	}
//
// 错误示例
//func main() {
//	doIt := func(arg int) interface{} {
//		var result *struct{} = nil
//		if arg > 0 {
//			result = &struct{}{}
//		}
//		return result
//	}
//
//	if res := doIt(-1); res != nil {
//		fmt.Println("Good result: ", res) // Good result:  <nil>
//		fmt.Printf("%T\n", res)           // *struct {}    // res 不是 nil，它的值为 nil
//		fmt.Printf("%v\n", res)           // <nil>
//	}
//}

// 正确示例
func main17() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil // 明确指明返回 nil
		}
		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res) // Bad result:  <nil>
	}
}
