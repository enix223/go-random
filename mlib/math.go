package mlib

import "math"

// LogBase log function with base
func LogBase(x, base float64) float64 {
	if base == 0 {
		return 0
	}
	return math.Log(x) / math.Log(base)
}
