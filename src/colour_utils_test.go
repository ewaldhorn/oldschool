package main

import "testing"

// ----------------------------------------------------------------------------
func Test_randomDelta(t *testing.T) {
	for i := range 5000 {
		result := randomDelta()

		if result != -1 && result != 1 {
			t.Errorf("test %d failed with %d.", i, result)
		}
	}
}

// ----------------------------------------------------------------------------
// updateColourValueWithinLimits(colour *uint8, change *int8, minColour uint8, maxColour uint8) {
func Test_updateColourValueWithinLimits(t *testing.T) {
	tests := []struct {
		name           string
		colour         uint8
		change         int8
		min            uint8
		max            uint8
		expectedColour uint8
		expectedChange int8
	}{
		{"one", 200, 1, 100, 250, 201, 1},
		{"two", 200, 1, 100, 200, 201, -1},
		{"three", 1, -1, 1, 50, 0, 1},
	}

	for _, test := range tests {
		updateColourValueWithinLimits(&test.colour, &test.change, test.min, test.max)

		if test.colour != test.expectedColour || test.change != test.expectedChange {
			t.Errorf("%s failed with colour %d, expected %d and change %d, expected %d", test.name, test.colour, test.expectedColour, test.change, test.expectedChange)
		}
	}
}
