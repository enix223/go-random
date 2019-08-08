# go-random

Random library implement by golang

## 0. Install

```
go get -u github.com/enix223/go-random
```

## 1. Collection

### 1.1 `Sample`
> Chooses k unique random elements from population with slice/array/string type

* signature

```golang
func Sample(population interface{}, k int) interface{}
```

* usage

```golang
import (
    "github.com/enix223/go-random/collection"
)

// call with string
// choose 10 items from alphabet letters
// r will be something like this: "AkuOMNBsqe"
r := collection.Sample(random.Alphabet, 10)

// call with slice
// choose 10 items from alphabet letters slice
// r will be something like this: []string{"A", "k", "u", "O", "M", "N", "B", "s", "q", "e"}
r := collection.Sample(random.AlphabetSlice, 10)

// call with array
// choose 10 items from alphabet letters
// r will be something like this: [3]rune{68, 65, 67}
alphabetArray := [10]rune{65, 66, 67, 68}
r := collection.Sample(alphabetArray, 3)
```

### 1.2 `SampleStringSlice`
> Chooses k unique random elements from population with string slice type

* signature

```golang
func SampleStringSlice(population []string, k int) []string
```

* usage

```golang
import (
    "github.com/enix223/go-random/collection"
)

// choose 10 items from alphabet letters slice
// r will be something like this: []string{"A", "k", "u", "O", "M", "N", "B", "s", "q", "e"}
r := collection.SampleStringSlice(random.AlphabetSlice, 10)
```

## 2. Color

Random functions for colorspace

### 2.1 `RandomRGB`
> get a random rgb color array

* signature

```golang
func RandomRGB() [3]uint8
```

* example

```golang
import (
    "github.com/enix223/go-random/color"
)

// get a random color, eg., []uint8{240, 255, 0}
rgb := color.RandomRGB()
```

### 2.2 `RandomRGBA`
> get a random rgb color array

* signature

```golang
func RandomRGBA() [4]uint8
```

* example

```golang
import (
    "github.com/enix223/go-random/color"
)

// get a random color with alpha channel, eg., []uint8{240, 255, 0, 124}
rgba := color.RandomRGBA()
```

### 2.3 `RandomRGBString`
> get a random rgb color in string format "#ffffff"

* signature

```golang
func RandomRGBString() string
```

* example

```golang
import (
    "github.com/enix223/go-random/color"
)

// get a random color in string format, eg., #ff91ab
rgb := color.RandomRGBString()
```

### 2.4 `RandomRGBAString`
> get a random rgb color with alpha channel in string format "#ffffffff"

* signature

```golang
func RandomRGBAString() string
```

* example

```golang
import (
    "github.com/enix223/go-random/color"
)

// get a random color with alpha channel in string format, eg., #ff91ab19
rgba := color.RandomRGBAString()
```
