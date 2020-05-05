package collection

import (
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/enix223/go-random/mlib"
)

// Punctuation punctation letters
const Punctuation = "!\"#$%&\\'()*+,-./:;<=>?@[]^_`{|}~"

// AlphabetLower lower case alphabet
const AlphabetLower = "abcdefghijklmnopqrstuvwxyz"

// AlphabetUpper upper case alphabet
const AlphabetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Alphabet english alphabet string collection
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// AlphabetDigits english alphabet string and digits collection
const AlphabetDigits = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Digits digits collection
const Digits = "0123456789"

// HexDigits hex digits collection
const HexDigits = "0123456789abcdefABCDEF"

// AlphabetSlice english alphabet slice
var AlphabetSlice = strings.Split(Alphabet, "")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Sample Chooses k unique random elements from a population slice/array/string
// population should be a slice, or it will panic
// the return value will have the same type as population
func Sample(population interface{}, k int) interface{} {
	popVal := reflect.ValueOf(population)
	if k > popVal.Len() {
		panic("Sample larger than population")
	}

	if popVal.Len() == 0 {
		panic("population should not be empty")
	}

	var n int
	var result reflect.Value
	switch popVal.Kind() {
	case reflect.Slice:
		n = popVal.Len()
		result = reflect.MakeSlice(popVal.Type(), k, k)
	case reflect.Array:
		n = popVal.Len()
		result = reflect.New(reflect.ArrayOf(k, popVal.Index(0).Type())).Elem()
	case reflect.String:
		res := SampleStringSlice(strings.Split(population.(string), ""), k)
		return strings.Join(res, "")
	default:
		panic("population should be type of array, slice or string")
	}

	// size of a small set minus size of an empty list
	var setsize float64 = 21
	if k > 5 {
		// table size for big sets
		setsize = math.Pow(4, math.Ceil(mlib.LogBase(float64(k)*3, 4)))
	}

	if n <= int(setsize) {
		// copy the population to sample item pool
		var poolVal reflect.Value
		// Copy the kth item of src to ith item of target
		// suppport slice/array/string
		var assign func(target, src *reflect.Value, i, k int)
		switch popVal.Kind() {
		case reflect.Slice:
			poolVal = reflect.MakeSlice(popVal.Type(), popVal.Len(), popVal.Len())
			reflect.Copy(poolVal, popVal)
			assign = func(target, src *reflect.Value, i, k int) {
				target.Index(i).Set(src.Index(k))
			}
		case reflect.Array:
			poolVal = reflect.New(reflect.ArrayOf(popVal.Len(), popVal.Index(0).Type())).Elem()
			reflect.Copy(poolVal, popVal)
			assign = func(target, src *reflect.Value, i, k int) {
				target.Index(i).Set(src.Index(k))
			}
		case reflect.String:
			poolVal = reflect.New(popVal.Type()).Elem()
			poolVal.SetString(popVal.String())
			assign = func(target, src *reflect.Value, i, k int) {
				t := []rune(target.String())
				s := []rune(src.String())
				t[i] = s[k]
				target.SetString(string(t))
			}
		}

		for i := 0; i < k; i++ {
			r := rand.Intn(n - i)
			assign(&result, &poolVal, i, r)
			// move the last elem to replace the selected item
			assign(&poolVal, &poolVal, r, n-i-1)
		}
	} else {
		selected := make(map[int]struct{})

		for i := 0; i < k; i++ {
			j := rand.Intn(n)
			for _, ok := selected[j]; ok; {
				j = rand.Intn(n)
			}
			result.Index(i).Set(popVal.Index(j))
			selected[j] = struct{}{}
		}
	}
	return result.Interface()
}

// SampleStringSlice Chooses k unique random elements from string slice population
func SampleStringSlice(population []string, k int) []string {
	if k > len(population) {
		panic("Sample larger than population")
	}

	if population == nil || len(population) == 0 {
		panic("population should not be empty")
	}

	result := make([]string, k)
	n := len(population)

	// size of a small set minus size of an empty list
	var setsize float64 = 21
	if k > 5 {
		// table size for big sets
		setsize = math.Pow(4, math.Ceil(mlib.LogBase(float64(k)*3, 4)))
	}
	if n <= int(setsize) {
		// copy the population to sample item pool
		pool := make([]string, len(population))
		copy(pool, population)

		for i := 0; i < k; i++ {
			k := rand.Intn(n - i)
			result[i] = pool[k]
			// move the last elem to replace the selected item
			pool[k] = pool[n-i-1]
		}
	} else {
		selected := make(map[int]struct{})

		for i := 0; i < k; i++ {
			j := rand.Intn(n)
			for _, ok := selected[j]; ok; {
				j = rand.Intn(n)
			}
			result = append(result, population[j])
			selected[j] = struct{}{}
		}
	}
	return result
}
