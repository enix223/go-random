package mlib

import "testing"

func TestMean(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	m := Mean(x)
	if m != 3 {
		t.Fatalf("exp: %v, got: %v", 3, m)
	}
}

func TestMeanWithEmptySlice(t *testing.T) {
	x := []int{}
	m := Mean(x)
	if m != 0 {
		t.Fatalf("exp: %v, got: %v", 0, m)
	}
}

func TestStandardDeviation(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	s := Std(x)
	if s-1.245 > 0.001 {
		t.Fatalf("exp: %v, got: %v", 1.245, s)
	}
}

func TestStandardDeviationWithEmptySlice(t *testing.T) {
	x := []int{}
	s := Std(x)
	if s != 0 {
		t.Fatalf("exp: %v, got: %v", 0, s)
	}
}

func TestVariance(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	v := Var(x)
	if v != 1.55 {
		t.Fatalf("exp: %v, got: %v", 1.55, v)
	}
}

func TestVarianceWithEmptySlice(t *testing.T) {
	x := []int{}
	v := Var(x)
	if v != 0 {
		t.Fatalf("exp: %v, got: %v", 0, v)
	}
}
