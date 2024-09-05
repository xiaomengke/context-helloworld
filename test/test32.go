package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main32() {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("pp", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "wwwwww"})
	})
	r.Any("pg", func(c *gin.Context) {
		c.String(200, "hellosss")
	})
	r.Routes()
	r.Run("localhost:8000")
	s := []int{1, 2, 3, 4, 5, 6}
	b := []int{4, 4, 5, 6, 8, 1, 2}
	sb := make([]int, len(s)+len(b))
	fmt.Println(s, b)
	copy(sb, s)
	copy(sb[len(s):], b)
	fmt.Println(sb)
}
