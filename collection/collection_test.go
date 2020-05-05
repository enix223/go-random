package collection

import (
	"reflect"
	"strings"
	"testing"

	"github.com/enix223/go-random/mlib"
)

func TestSampleSlice(t *testing.T) {
	k := 10
	s := SampleStringSlice(AlphabetSlice, k)
	if len(s) != k {
		t.Fatalf("exp: %v, got: %v", k, len(s))
	}

	set := make(map[string]struct{})
	for _, i := range s {
		set[i] = struct{}{}
	}

	if len(set) != k {
		t.Fatalf("exp: %v, got: %v", k, len(set))
	}
}

func TestSampleSliceKLargerThanPopulation(t *testing.T) {
	k := 1000
	defer func() {
		e := "Sample larger than population"
		if err := recover(); err != e {
			t.Fatalf("exp :%v, got: %v", e, err)
		}
	}()
	SampleStringSlice(AlphabetSlice, k)
}

func TestSampleSliceRandom(t *testing.T) {
	stat := make([]int, len(AlphabetSlice))
	k := 10
	for i := 0; i < len(AlphabetSlice); i++ {
		s := SampleStringSlice(AlphabetSlice, k)
		for _, i := range s {
			stat[strings.Index(Alphabet, i)]++
		}
	}

	mean := mlib.Mean(stat)
	std := mlib.Std(stat)
	if std > mean {
		t.Fatalf("exp: %v, got: %v", "small std", std)
	}
}

func TestSampleWithSlice(t *testing.T) {
	k := 10
	r := Sample(AlphabetSlice, k)
	s, ok := r.([]string)
	if !ok {
		t.Fatalf("exp: %v, got: %v", "a slice of string", r)
	}

	if len(s) != k {
		t.Fatalf("exp: %v, got: %v", k, len(s))
	}

	set := make(map[string]struct{})
	for _, i := range s {
		set[i] = struct{}{}
	}

	if len(set) != k {
		t.Fatalf("exp: %v, got: %v", k, len(set))
	}
}

func TestSampleWithString(t *testing.T) {
	k := 5
	r := Sample(Alphabet, k)
	s, ok := r.(string)
	if !ok {
		t.Fatalf("exp: %v, got: %v", "a string", r)
	}

	if len(s) != k {
		t.Fatalf("exp: %v, got: %v", k, len(s))
	}

	set := make(map[rune]struct{})
	for _, i := range s {
		set[i] = struct{}{}
	}

	if len(set) != k {
		t.Fatalf("exp: %v, got: %v", k, len(set))
	}
}

func TestSampleWithArray(t *testing.T) {
	k := 10
	var population [52]rune
	for i, r := range Alphabet {
		population[i] = r
	}

	r := Sample(population, k)
	s, ok := r.([10]rune)
	if !ok {
		tt := reflect.TypeOf(r)
		t.Fatalf("exp: %v, got: %v", "an rune array", tt)
	}

	if len(s) != k {
		t.Fatalf("exp: %v, got: %v", k, len(s))
	}

	set := make(map[rune]struct{})
	for _, i := range s {
		set[i] = struct{}{}
	}

	if len(set) != k {
		t.Fatalf("exp: %v, got: %v", k, len(set))
	}
}

// Benchmark

// Compare the performance between generic Sample and SampleStringSlice version
func BenchmarkSample(b *testing.B) {
	k := 10

	b.Run("Sample", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			Sample(AlphabetSlice, k)
		}
	})

	b.Run("SampleStringSlice", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			SampleStringSlice(AlphabetSlice, k)
		}
	})
}
