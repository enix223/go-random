package color

import (
	"fmt"
	"math/rand"
)

// RandomRGB get a random rgb color array
func RandomRGB() [3]uint8 {
	var rgb [3]uint8
	rgb[0] = uint8(rand.Intn(256))
	rgb[1] = uint8(rand.Intn(256))
	rgb[2] = uint8(rand.Intn(256))
	return rgb
}

// RandomRGBA get a random rgb color array
func RandomRGBA() [4]uint8 {
	var rgba [4]uint8
	rgba[0] = uint8(rand.Intn(256))
	rgba[1] = uint8(rand.Intn(256))
	rgba[2] = uint8(rand.Intn(256))
	rgba[3] = uint8(rand.Intn(256))
	return rgba
}

// RandomRGBString get a random rgb color in string format "#ffffff"
func RandomRGBString() string {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// RandomRGBAString get a random rgb color with alpha channel in string format "#ffffffff"
func RandomRGBAString() string {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	a := uint8(rand.Intn(256))
	return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}
