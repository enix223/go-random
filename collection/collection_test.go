package collection

import (
	"reflect"
	"strings"
	"testing"

	"github.com/enix223/go-random/mlib"
)

func TestSampleWithInvalidPopulation(t *testing.T) {
	k := 10
	defer func() {
		e := "population should be type of array, slice or string"
		if err := recover(); err != e {
			t.Fatalf("exp :%v, got: %v", e, err)
		}
	}()
	Sample(123, k)
}

func TestSampleKLargerThanPopulation(t *testing.T) {
	k := 1000
	defer func() {
		e := "Sample larger than population"
		if err := recover(); err != e {
			t.Fatalf("exp :%v, got: %v", e, err)
		}
	}()
	Sample(AlphabetSlice, k)
}

func TestSampleWithEmptyPopulation(t *testing.T) {
	k := 0
	defer func() {
		e := "population should not be empty"
		if err := recover(); err != e {
			t.Fatalf("exp :%v, got: %v", e, err)
		}
	}()
	Sample([]int{}, k)
}

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

func TestSampleSliceWithEmptyPopulation(t *testing.T) {
	k := 0
	defer func() {
		e := "population should not be empty"
		if err := recover(); err != e {
			t.Fatalf("exp :%v, got: %v", e, err)
		}
	}()
	SampleStringSlice([]string{}, k)
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

func TestRandomWithSlice(t *testing.T) {
	fn := func(p []string, k int) {
		r := Random(p, k)
		res, ok := r.([]string)
		if !ok {
			t.Fatalf("result is not []string")
		}
		exp := k
		if len(res) != exp {
			t.Fatalf("exp: %v, got: %v", exp, len(res))
		}

		for _, i := range res {
			found := false
			for _, e := range p {
				if i == e {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("element not found in population")
			}
		}
	}

	fn([]string{"1", "2", "3"}, 3)
	fn([]string{"1", "2", "3"}, 10)
	fn([]string{"1"}, 10)
}

func TestRandomWithArray(t *testing.T) {
	p1 := [3]string{"1", "2", "3"}
	r := Random(p1, 10)
	res1, ok := r.([10]string)
	if !ok {
		t.Fatalf("result is not []string")
	}
	for _, i := range res1 {
		found := false
		for _, e := range p1 {
			if i == e {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("element not found in population")
		}
	}

	p2 := [1]string{"1"}
	r = Random(p2, 3)
	res2, ok := r.([3]string)
	if !ok {
		t.Fatalf("result is not []string")
	}
	for _, i := range res2 {
		found := false
		for _, e := range p2 {
			if i == e {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("element not found in population")
		}
	}
}

func TestRandomWithString(t *testing.T) {
	fn := func(p string, k int) {
		r := Random(p, k)
		res1, ok := r.(string)
		if !ok {
			t.Fatalf("result is not string")
		}
		if len(res1) != k {
			t.Fatalf("exp length: %v, got: %v", k, len(res1))
		}
		for _, i := range res1 {
			found := false
			for _, e := range p {
				if i == e {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("element not found in population")
			}
		}
	}
	fn("1234", 10)
	fn("1234", 1)
	fn("1234", 4)
}

func TestRandomStringSliceWithCorrectParam(t *testing.T) {
	p := []string{"1", "2", "3"}
	r := RandomStringSlice(p, 3)
	if len(r) != 3 {
		t.Fatalf("exp: %v, got: %v", 3, len(r))
	}

	p = []string{"1"}
	r = RandomStringSlice(p, 3)
	expl := 3
	if len(r) != expl {
		t.Fatalf("exp: %v, got: %v", expl, len(r))
	}
	expr := []string{"1", "1", "1"}
	if !reflect.DeepEqual(r, expr) {
		t.Fatalf("exp: %v, got: %v", expr, r)
	}
}

func TestRandomWithBadParam(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("exp: population should be type of array, slice or string, got: nil")
		}
	}()

	Random(123, 10)
}

func TestRandomWithEmptyParam(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("exp: population should not be empty, got: nil")
		}
	}()

	Random("", 10)
}

func TestRandomWithEmptySlice(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("exp: population should not be empty, got: nil")
		}
	}()

	Random([]int{}, 10)
}

func TestRandomStringSliceWithBadParam(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("exp: population should not be empty, got: nil")
		}
	}()

	RandomStringSlice([]string{}, 10)
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
