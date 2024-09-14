package swap

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		input []interface{}
		i     int
		j     int
		want  []interface{}
	}{
		{[]interface{}{1, 2}, 0, 1, []interface{}{2, 1}},
		{[]interface{}{1, 'a', "aa"}, 0, 2, []interface{}{"aa", 'a', 1}},
	}
	for i, tt := range tests {
		if err := Swap(tt.input, tt.i, tt.j); err != nil {
			t.Error(err)
		}

		if !IsSameSlice(tt.input, tt.want) {
			t.Errorf("%v. got %v, want %v", i, tt.input, tt.want)
		}
	}
}
