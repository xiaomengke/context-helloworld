package shendun_exam

import (
	"fmt"
	"math/rand"
)

const IdNum = 1000
const OrderNum = 1000000

var idSlice []int

func generateId() []int {
	generatedIds := make(map[int]bool)
	userIds := make([]int, IdNum)
	for i := 0; i < IdNum; {
		id := rand.Intn(10000000)
		if _, exists := generatedIds[id]; !exists {
			generatedIds[id] = true
			userIds[i] = id
			i++
		}
	}
	return userIds
}

func GenerateOrder() {
	idSlice = generateId()
	fmt.Println(idSlice)
}

func insertToSql() {

}

func generateWeight() (re float64) {
	//有各事件出现具体概率的话，根据累积概率分布就可以获取发生事件。
	// 计算累积概率
	probabilities := make([]float64, 100)
	sum := 0.0
	for i := 0; i < 100; i++ {
		sum += 1.0 / (float64(i+1) * 1.0)
		probabilities[i] = sum
	}
	// 生成随机数
	r := rand.Float64() * sum
	// 找到对应的整数,然后取其范围内的double
	for i, p := range probabilities {
		if r <= p {
			re = float64(i+1) - rand.Float64()
			break
		}
	}
	return
}

func QueryOrderById() {

}
