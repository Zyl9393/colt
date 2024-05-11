package colt

import "image/color"

// https://en.wikipedia.org/wiki/Relative_luminance
const luminanceRed float32 = 0.2126
const luminanceGreen float32 = 0.7152
const luminanceBlue float32 = 0.0722

// SRGB represents an sRGB color where each color component is encoded in a byte.
type SRGB [3]uint8

// SRGBA represents an sRGB color with alpha where each color component is encoded in a byte.
type SRGBA [4]uint8

// RGB represents a linear RGB color.
type RGB [3]float32

// RGBA represents a linear RGB color with alpha.
type RGBA [4]float32

// Std returns c as a color.RGBA, assuming it is alpha-premultiplied.
func (c SRGB) Std() color.RGBA {
	return color.RGBA{c[0], c[1], c[2], 1}
}

// StdN returns c as a color.NRGBA, assuming it is not alpha-premultiplied.
func (c SRGB) StdN() color.NRGBA {
	return color.NRGBA{c[0], c[1], c[2], 1}
}

// Std returns c as a color.RGBA, assuming it is alpha-premultiplied.
func (c SRGBA) Std() color.RGBA {
	return color.RGBA{c[0], c[1], c[2], c[3]}
}

// StdN returns c as a color.NRGBA, assuming it is not alpha-premultiplied.
func (c SRGBA) StdN() color.NRGBA {
	return color.NRGBA{c[0], c[1], c[2], c[3]}
}

// Std lossily returns c as a color.RGBA64, assuming it is alpha-premultiplied.
func (c RGB) Std() color.RGBA64 {
	return color.RGBA64{Standard16(c[0]), Standard16(c[1]), Standard16(c[2]), 1}
}

// StdN lossily returns c as a color.NRGBA64, assuming it is not alpha-premultiplied.
func (c RGB) StdN() color.NRGBA64 {
	return color.NRGBA64{Standard16(c[0]), Standard16(c[1]), Standard16(c[2]), 1}
}

// Std lossily returns c as a color.RGBA64, assuming it is alpha-premultiplied.
func (c RGBA) Std() color.RGBA64 {
	return color.RGBA64{Standard16(c[0]), Standard16(c[1]), Standard16(c[2]), uint16(c[3]*0xffff + 0.5)}
}

// StdN lossily returns c as a color.NRGBA64, assuming it is not alpha-premultiplied.
func (c RGBA) StdN() color.NRGBA64 {
	return color.NRGBA64{Standard16(c[0]), Standard16(c[1]), Standard16(c[2]), uint16(c[3]*0xffff + 0.5)}
}

// Linear returns the linear RGB representation of this sRGB-encoded color.
func (c SRGB) Linear() RGB {
	return RGB{Linearb(c[0]), Linearb(c[1]), Linearb(c[2])}
}

// Linear returns the linear RGB representation of this sRGB-encoded color with alpha.
func (c SRGBA) Linear() RGBA {
	return RGBA{Linearb(c[0]), Linearb(c[1]), Linearb(c[2]), ByteDecoded(c[3])}
}

// Standard returns the sRGB-encoded representation of this linear RGB color.
func (c RGB) Standard() SRGB {
	return SRGB{Standardb(c[0]), Standardb(c[1]), Standardb(c[2])}
}

// Standard returns the sRGB-encoded representation of this linear RGB color with alpha.
func (c RGBA) Standard() SRGBA {
	return SRGBA{Standardb(c[0]), Standardb(c[1]), Standardb(c[2]), ByteEncoded(c[3])}
}

// Alpha returns SRGBA{c[0], c[1], c[2], alpha}.
func (c SRGB) Alpha(alpha uint8) SRGBA {
	return SRGBA{c[0], c[1], c[2], alpha}
}

// SRGB returns SRGB{c[0], c[1], c[2]}.
func (c SRGBA) SRGB(alpha uint8) SRGB {
	return SRGB{c[0], c[1], c[2]}
}

// Luminance returns the luminance of the color using standard luminance values.
func (c RGB) Luminance() float32 {
	return c[0]*luminanceRed + c[1]*luminanceGreen + c[2]*luminanceBlue
}

// Luminance returns the luminance of the color using standard luminance values, multiplied by its alpha value.
func (c RGBA) Luminance() float32 {
	return (c[0]*luminanceRed + c[1]*luminanceGreen + c[2]*luminanceBlue) * c[3]
}

// LuminanceCustom returns the luminance of the color using given luminance values.
func (c RGB) LuminanceCustom(luminanceRed, luminanceGreen, luminanceBlue float32) float32 {
	return c[0]*luminanceRed + c[1]*luminanceGreen + c[2]*luminanceBlue
}

// LuminanceCustom returns the luminance of the color using given luminance values, multiplied by its alpha value.
func (c RGBA) LuminanceCustom(luminanceRed, luminanceGreen, luminanceBlue float32) float32 {
	return (c[0]*luminanceRed + c[1]*luminanceGreen + c[2]*luminanceBlue) * c[3]
}

// Blend blends c with other*amount.
func (c RGB) Blend(other RGB, amount float32) RGB {
	if amount < 0 {
		amount = 0
	} else if amount > 1 {
		amount = 1
	}
	inverse := 1 - amount
	return RGB{c[0]*inverse + other[0]*amount, c[1]*inverse + other[1]*amount, c[2]*inverse + other[2]*amount}
}

// Blend blends c with other*amount.
func (c RGBA) Blend(other RGBA, amount float32) RGBA {
	if amount < 0 {
		amount = 0
	} else if amount > 1 {
		amount = 1
	}
	inverse := 1 - amount
	return RGBA{c[0]*inverse + other[0]*amount, c[1]*inverse + other[1]*amount, c[2]*inverse + other[2]*amount, c[3]*inverse + other[3]*amount}
}

// Mul multiplies each component of c with amount.
func (c RGB) Mul(amount float32) RGB {
	return RGB{c[0] * amount, c[1] * amount, c[2] * amount}
}

// Mul multiplies each component of c with amount.
func (c RGBA) Mul(amount float32) RGBA {
	return RGBA{c[0] * amount, c[1] * amount, c[2] * amount, c[3] * amount}
}

// Add adds other to c.
func (c RGB) Add(other RGB) RGB {
	return RGB{c[0] + other[0], c[1] + other[1], c[2] + other[2]}
}

// Add adds other to c.
func (c RGBA) Add(other RGBA) RGBA {
	return RGBA{c[0] + other[0], c[1] + other[1], c[2] + other[2], c[3] + other[3]}
}

// Sub subtracts other from c.
func (c RGB) Sub(other RGB) RGB {
	return RGB{c[0] - other[0], c[1] - other[1], c[2] - other[2]}
}

// Sub subtracts other from c.
func (c RGBA) Sub(other RGBA) RGBA {
	return RGBA{c[0] - other[0], c[1] - other[1], c[2] - other[2], c[3] - other[3]}
}

// Alpha returns RGBA{c[0], c[1], c[2], alpha}.
func (c RGB) Alpha(alpha float32) RGBA {
	return RGBA{c[0], c[1], c[2], alpha}
}

// RGB returns RGB{c[0], c[1], c[2]}.
func (c RGBA) RGB() RGB {
	return RGB{c[0], c[1], c[2]}
}

func (c *RGB) clamp() {
	for i := 0; i < 3; i++ {
		if c[i] < 0 {
			c[i] = 0
		} else if c[i] > 1 {
			c[i] = 1
		}
	}
}

func (c *RGBA) clamp() {
	for i := 0; i < 4; i++ {
		if c[i] < 0 {
			c[i] = 0
		} else if c[i] > 1 {
			c[i] = 1
		}
	}
}

// Clamp clamps c's components to range [0;1].
func (c RGB) Clamp() RGB {
	for i := 0; i < 3; i++ {
		if c[i] < 0 {
			c[i] = 0
		} else if c[i] > 1 {
			c[i] = 1
		}
	}
	return c
}

// Clamp clamps c's components to range [0;1].
func (c RGBA) Clamp() RGBA {
	for i := 0; i < 4; i++ {
		if c[i] < 0 {
			c[i] = 0
		} else if c[i] > 1 {
			c[i] = 1
		}
	}
	return c
}

// Premultiplied multiplies the R, G and B component of this color with its alpha component,
// such that in a normal blending scenario the new color is then given by source + dest * (1 - c[3])
// instead of source * c[3] + dest * (1 - c[3]), in hopes of easing computational load.
//
// It is the caller's responsibility to track whether this RGBA color is currently premulitplied or not.
func (c RGBA) Premultiplied() RGBA {
	c[0] *= c[3]
	c[1] *= c[3]
	c[2] *= c[3]
	return c
}

// UnPremultiplied reverts a prior multiplication of the R, G and B components of this color
// with its alpha value.
//
// It is the caller's responsibility to track whether this RGBA color is currently premulitplied or not.
func (c RGBA) UnPremultiplied() RGBA {
	if c[3] == 0 {
		return c
	}
	c[0] /= c[3]
	c[1] /= c[3]
	c[2] /= c[3]
	return c
}

// ByteEncoded lossily encodes the range [0.0; 1.0] of a float32 value into a byte using range [0; 255] and returns it.
func ByteEncoded(v float32) uint8 {
	if v >= 1 {
		return 255
	}
	if v <= 0 {
		return 0
	}
	return uint8(float32(v)*255 + 0.5)
}

// ByteDecoded decodes the range [0; 255] of a byte value into a float32 using range [0.0; 1.0] and returns it.
func ByteDecoded(v uint8) float32 {
	return byteDecoded[v]
}
