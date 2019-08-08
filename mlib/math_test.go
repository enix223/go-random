package mlib

import (
	"testing"
)

func TestLogBase(t *testing.T) {
	x := LogBase(4, 2)
	if x != 2 {
		t.Fatalf("exp: %v, got: %v", 2, x)
	}
}

func TestLogBaseZero(t *testing.T) {
	x := LogBase(0, 0)
	if x != 0 {
		t.Fatalf("exp: %v, got: %v", 0, x)
	}
}
