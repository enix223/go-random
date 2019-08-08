package mlib

import "math"

// Var calculate the variance of the given population set
// \operatorname {Var} (X)=\operatorname {E} \left[(X-\mu )^{2}\right].
func Var(population []int) float64 {
	if len(population) == 0 {
		return 0
	}

	u := Mean(population)
	var t float64
	for _, i := range population {
		t += math.Pow(2, float64(i)-u)
	}
	t /= float64(len(population))
	return t
}

// Std calculate the standard deviation of th given population set
// {\displaystyle s={\sqrt {{\frac {1}{N-1}}\sum _{i=1}^{N}(x_{i}-{\bar {x}})^{2}}},}
func Std(population []int) float64 {
	if len(population) == 0 {
		return 0
	}
	return math.Sqrt(Var(population))
}

// Mean calculate the mean of the given population set
func Mean(population []int) float64 {
	if len(population) == 0 {
		return 0
	}

	var u float64
	var t int
	for _, i := range population {
		t += i
	}
	u = float64(t) / float64(len(population))
	return u
}
