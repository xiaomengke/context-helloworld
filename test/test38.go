package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main38() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 计算累积概率
	probabilities := make([]float64, 100)
	sum := 0.0
	for i := 0; i < 100; i++ {
		sum += 1.0 / (float64(i+1) * 1.0)
		probabilities[i] = sum
	}
	fmt.Println(probabilities)

	// 生成随机数
	r := rand.Float64() * sum
	fmt.Println(r)
	// 找到对应的整数
	for i, p := range probabilities {
		if r <= p {
			fmt.Println(i + 1)
			break
		}
	}
}
