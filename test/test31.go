package main

import (
	"fmt"
	"reflect"
)

func main31() {
	num1 := 666
	fmt.Println("num1 原值：", num1)
	func4Reflect(&num1)
	fmt.Println("num1 修改后的值：", num1)

	num2 := 0.5
	fmt.Println("num2 原值：", num2)
	func4Reflect(&num2)
	fmt.Println("num2 修改后的值：", num2)

	str := "go"
	fmt.Println("str 原值：", str)
	func4Reflect(&str)
	fmt.Println("str 修改后的值：", str)
}

func func4Reflect(data any) {
	typ := reflect.TypeOf(data)
	fmt.Println(typ)
	val := reflect.ValueOf(data)
	fmt.Println(val)
	fmt.Println(val.Elem())
	fmt.Println(val.Elem().Kind())
	fmt.Println(val.Elem().Type())
	switch typ.Elem().Kind() {
	case reflect.Int:
		val.Elem().SetInt(888)
	case reflect.Float64:
		val.Elem().SetFloat(3.14)
	case reflect.String:
		val.Elem().SetString("Golang")
	}
}

type User22 struct {
	Name string
	Age  int
}

func main312() {
	user := User22{
		Name: "cmy",
		Age:  18,
	}
	func4Reflect2(user)
}

func func4Reflect2(data any) {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	// 获取结构体字段的数量
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		fmt.Println("字段名称：", typ.Field(i).Name)
		fmt.Println("字段类型：", typ.Field(i).Type)
		fmt.Println("字段值：", val.Field(i).Interface())
		fmt.Printf("%T\n", val.Field(i).Interface())
		fmt.Println("----------------------------")
	}
}
