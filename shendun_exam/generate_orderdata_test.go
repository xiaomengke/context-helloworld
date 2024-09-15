package shendun_exam

import "testing"

func TestGenerateOrder(t *testing.T) {
	GenerateOrder()
}

func TestQueryOrder(t *testing.T) {
	st, err := QueryOrderById(1554)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(st)
}
