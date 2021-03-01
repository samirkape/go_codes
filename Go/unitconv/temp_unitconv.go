package unitconv

//TKtoC Convert Kelvin to Celsius
func TKtoC(k Kelvin) Celsius {
	return Celsius(k - ZeroK)
}

//TFtoC Convert Fahrenheit to Celsius
func TFtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//TKtoF Convert Kelvin to Fahrenheit
func TKtoF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-ZeroK)*9/5 + 32)
}

//TCtoF Convert Celsius to Fahrenheit
func TCtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//TCtoK Convert Celsius to Kelvin
func TCtoK(c Celsius) Kelvin {
	return Kelvin(c - ZeroK)
}

//TFtoK Convert Fahrenheit to Kelvin
func TFtoK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}
