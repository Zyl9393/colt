# 3.2.0 (2024-07-07)
* Make luminance constants public.

# 3.1.1 (2024-05-13)
* Fixed implementations of `SRGB.Std()`, `SRGB.StdN()`, `RGB.Std()` and `RGB.StdN()` using wrong alpha.

# 3.1.0 (2024-05-13)
* Added `RGB.SubC()`.
* Fixed API documentation of `RGBA.SubC()` and `RGBA.Sub3()`.

# 3.0.0 (2024-05-12)
* **BREAKING:** Removed unused `alpha` parameter from `SRGBA.SRGB()`.
* **BREAKING:** Renamed `SRGBA.SRGB()` to `SRGBA.NoAlpha()` and `RGBA.RGB()` to `RGBA.NoAlpha()`.

# 2.0.1 (2024-05-12)
* Added `/v2` to module name.

# 2.0.0 (2024-05-12)
* **BREAKING:** Renamed `RGB.Multiply()` to `RGB.Mul()` and `RGBA.Multiply()` to `RGBA.Mul()`.
* Added `RGB.Sub()` and `RGBA.Sub()`.
* Added `SRGB.Alpha()`, `SRGBA.SRGB()` `RGB.Alpha()` and `RGBA.RGB()`.
* Added `Standard16()` and `Linear16()`, `Linear16AllNew()` and `Standard16AllNew()`.
* Added `SRGB.Std()`, `SRGB.StdN()`, `SRGBA.Std()`, `SRGBA.StdN()`, `RGB.Std()`, `RGB.StdN()`, `RGBA.Std()` and `RGBA.StdN()`.
* Added `RGBA.AddC()`, `RGBA.Add3()`, `RGBA.SubC()` and `RGBA.Sub3()`.
