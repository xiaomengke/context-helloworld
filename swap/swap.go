package swap

import (
	"errors"
)

// Swap exchanges s[i] and s[j].
func Swap(s []interface{}, i, j int) error {
	if s == nil {
		return errors.New("slice can't be nil")
	}
	if (i < 0 || i >= len(s)) || (j < 0 || j >= len(s)) {
		return errors.New("illegal index")
	}

	s[i], s[j] = s[j], s[i]

	return nil
}

// IsSameSlice determines two slice is it the same.
func IsSameSlice(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
