package main

import "testing"

func Test_clampFloat32(t *testing.T) {
	tests := []struct {
		name     string
		value    float32
		min      float32
		max      float32
		expected float32
	}{
		{"zero", -19, 0, 10, 0},
		{"one", 1, 1, 5, 1},
		{"five", 5, 1, 8, 5},
		{"eleven", 18, 5, 11, 11},
	}

	for _, test := range tests {
		result := clampFloat32(test.value, test.min, test.max)

		if result != test.expected {
			t.Errorf("%s failed with %f instead of %f\n", test.name, result, test.expected)
		}
	}
}
