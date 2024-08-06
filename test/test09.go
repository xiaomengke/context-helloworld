package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	if err := json.Unmarshal(data, &result); err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Printf("%T\n", result["status"])    // float64
//	var status = result["status"].(float64) // 类型断言错误
//	fmt.Println("Status value: ", status)
//}

//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result struct {
//		Status int `json:"status"` //使用 struct 类型将你需要的数据映射为数值型
//	}
//
//	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
//	if err != nil {
//		fmt.Println("err")
//	}
//	fmt.Printf("Result: %+v", result)
//}

// 状态名称可能是 int 也可能是 string，指定为 json.RawMessage 类型
func main9() {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}

}
