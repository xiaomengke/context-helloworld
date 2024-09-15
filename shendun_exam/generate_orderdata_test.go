package shendun_exam

import (
	"fmt"
	"testing"
)

func TestGenerateOrder(t *testing.T) {
	GenerateOrder()
}

func TestQueryOrder(t *testing.T) {
	st, err := QueryOrderById(574)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(st)
}
func TestQueryOrderById(t *testing.T) {
	var id int
	var err error
	id, err = fmt.Scanf("%d", &id)
	if err != nil {
		t.Error(err.Error())
	}
	st, err := QueryOrderById(id)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(st)
}
