package main

// ----------------------------------------------------------------------------
// Clamps the value to the min/max
func clampFloat32(value, min, max float32) float32 {
	if min > max {
		panic("min must less than max")
	}

	if value <= min {
		return min
	}

	if max <= value {
		return max
	}

	return value
}
