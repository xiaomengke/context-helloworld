package shendun_exam

import (
	"fmt"
	"testing"
)

func TestGetCourierFee(t *testing.T) {
	testList := []float64{0.4, 1, 1.2, 2, 2.3, 98, 98.6, 99, 99.0, 99.2, 99.9, 100, 100.1, 100.9, 101, 108}
	for _, j := range testList {
		//fmt.Println(int(j))
		fmt.Println(GetCourierFee(j))
	}
}
