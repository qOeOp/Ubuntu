package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvins {
	return Kelvins(c + 273.15)
}

func FToK(f Fahrenheit) Kelvins {
	return CToK(FToC(f))
}

func KToF(k Kelvins) Fahrenheit {
	return CToF(KToC(k))
}
