package color

import "testing"

func TestRandomRGB(t *testing.T) {
	for i := 0; i < 100; i++ {
		rgb := RandomRGB()
		if rgb[0] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgb[0])
		}
		if rgb[1] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgb[1])
		}
		if rgb[2] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgb[2])
		}
	}
}

func TestRandomRGBString(t *testing.T) {
	for i := 0; i < 100; i++ {
		rgb := RandomRGBString()
		if rgb > "#ffffff" {
			t.Fatalf("exp: %v, exp: %v", "color less than #ffffff", rgb[2])
		}
	}
}

func TestRandomRGBA(t *testing.T) {
	for i := 0; i < 100; i++ {
		rgba := RandomRGBA()
		if rgba[0] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgba[0])
		}
		if rgba[1] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgba[1])
		}
		if rgba[2] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgba[2])
		}
		if rgba[3] > 255 {
			t.Fatalf("exp: %v, exp: %v", "less than 255", rgba[3])
		}
	}
}

func TestRandomRGBAString(t *testing.T) {
	for i := 0; i < 100; i++ {
		rgba := RandomRGBAString()
		if rgba > "#ffffffff" {
			t.Fatalf("exp: %v, exp: %v", "color less than #ffffffff", rgba)
		}
	}
}
