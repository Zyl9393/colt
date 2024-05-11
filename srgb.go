package colt

import "math"

// https://en.wikipedia.org/wiki/SRGB#Transformation

var byteDecoded [256]float32
var standardToLinear [256]float32

func init() {
	for i := 0; i < 256; i++ {
		byteDecoded[i] = float32(i) / 255
		standardToLinear[i] = Linear(byteDecoded[i])
	}
	standardToLinear[0] = 0
	standardToLinear[255] = 1
}

// Standard returns the sRGB color space value in range [0.0-1.0] for v, assuming v is in linear RGB in range [0.0-1.0].
func Standard(v float32) float32 {
	if v <= 0.0031308 {
		return v * 12.92
	}
	return float32(1.055*math.Pow(float64(v), 1.0/2.4) - 0.055)
}

// Standardb returns the sRGB color space value in range [0-255] for v, assuming v is in linear RGB in range [0.0-1.0].
func Standardb(v float32) uint8 {
	if v >= 1 {
		return 255
	}
	if v <= 0 {
		return 0
	}
	return uint8(Standard(v)*255 + 0.5)
}

// Standard16 returns the sRGB color space value in range [0-65535] for v, assuming v is in linear RGB in range [0.0-1.0].
func Standard16(v float32) uint16 {
	if v >= 1 {
		return 255
	}
	if v <= 0 {
		return 0
	}
	return uint16(Standard(v)*0xffff + 0.5)
}

// Linear returns the linear RGB color space value in range [0.0-1.0] for v, assuming v is in sRGB in range [0.0-1.0].
func Linear(v float32) float32 {
	if v <= 0.04045 {
		return v * (1.0 / 12.92)
	}
	return float32(math.Pow((float64(v)+0.055)/1.055, 2.4))
}

// Linearb returns the linear RGB color space value in range [0.0-1.0] for b, assuming b is in sRGB in range [0-255].
func Linearb(b uint8) float32 {
	return standardToLinear[b]
}

// Linear16 returns the linear RGB color space value in range [0.0-1.0] for v, assuming v is in sRGB in range [0-65535].
func Linear16(v uint16) float32 {
	return Linear(float32(v) / 0xffff)
}
