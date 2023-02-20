package colt

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

// Multiply multiplies each component of c with amount.
func (c RGB) Multiply(amount float32) RGB {
	return RGB{c[0] * amount, c[1] * amount, c[2] * amount}
}

// Multiply multiplies each component of c with amount.
func (c RGBA) Multiply(amount float32) RGBA {
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
