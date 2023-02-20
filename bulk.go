package colt

// LinearAll linearizes each element in values assuming it is an sRGB-encoded color component.
func LinearAll(values []float32) {
	for i, v := range values {
		values[i] = Linear(v)
	}
}

// LinearAllNew returns a new slice containing the linearized elements of values assuming they are sRGB-encoded color components, or nil if values is empty.
func LinearAllNew(values []float32) []float32 {
	if len(values) == 0 {
		return nil
	}
	linear := make([]float32, len(values), len(values))
	for i, v := range values {
		linear[i] = Linear(v)
	}
	return linear
}

// LinearbAllNew returns a new slice containing the linearized elements of values assuming they are sRGB-encoded color components, or nil if values is empty.
func LinearbAllNew(values []uint8) []float32 {
	if len(values) == 0 {
		return nil
	}
	linear := make([]float32, len(values), len(values))
	for i, v := range values {
		linear[i] = Linearb(v)
	}
	return linear
}

// StandardAll encodes each element in values in sRGB assuming it is an linear RGB color component.
func StandardAll(values []float32) {
	for i, v := range values {
		values[i] = Standard(v)
	}
}

// StandardAllNew returns a new slice containing the sRGB-encoded elements of values assuming they are linear RGB color components, or nil if values is empty.
func StandardAllNew(values []float32) []float32 {
	if len(values) == 0 {
		return nil
	}
	standard := make([]float32, len(values), len(values))
	for i, v := range values {
		standard[i] = Standard(v)
	}
	return standard
}

// StandardbAllNew returns a new slice containing the sRGB-encoded elements of values assuming they are linear RGB color components, or nil if values is empty.
func StandardbAllNew(values []float32) []uint8 {
	if len(values) == 0 {
		return nil
	}
	standard := make([]uint8, len(values), len(values))
	for i, v := range values {
		standard[i] = Standardb(v)
	}
	return standard
}
