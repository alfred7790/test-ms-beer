package utilities

import "testing"

func TestToFixed(t *testing.T) {
	tests := []struct {
		number    float64
		expected  float64
		precision int
	}{
		{number: 12.55764, expected: 12.56, precision: 2},
		{number: .99999, expected: 1.00, precision: 2},
		{number: .0065, expected: .01, precision: 2},
		{number: 0.1234, expected: 0.12, precision: 2},
		{number: 0.1234, expected: 0.1234, precision: 4},
		{number: 0.1234, expected: 0.123, precision: 3},
		{number: 0.1234, expected: 0.1, precision: 1},
		{number: 0.1234, expected: 0, precision: 0},
		{number: 0.1234, expected: 0, precision: 0},
		{number: 546, expected: 546, precision: 0},
	}

	for _, v := range tests {
		if v.expected != ToFixed(v.number, v.precision) {
			t.Errorf("expected: %v actual:%v", v.expected, ToFixed(v.number, v.precision))
		}
	}
}
