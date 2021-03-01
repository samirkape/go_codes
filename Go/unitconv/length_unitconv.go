//This file has functions for Length conversion

package unitconv

//LCtoI Convert CM to Inches
func LCtoI(c CM) Inch {
	return Inch(c / OneIN_CM)
}

//LCtoM Convert to
func LCtoM(c CM) Meter {
	return Meter(c / OneM_CM)
}

//LCtoF Convert CM to Feet
func LCtoF(c CM) Feet {
	return Feet(c / OneF_CM)
}

//LItoC Convert Inch to CM
func LItoC(i Inch) CM {
	return CM(i * OneIN_CM)
}

//LItoM Convert Inch to Meter
func LItoM(i Inch) Meter {
	return Meter(LItoC(i) / OneM_CM)
}

//LItoF Convert Inch to Feet
func LItoF(i Inch) Feet {
	return Feet(LItoC(i) / OneF_CM)
}

//LMtoC Convert Meter to CM
func LMtoC(m Meter) CM {
	return CM(m * OneM_CM)
}

//LMtoI Convert Meter to Inch
func LMtoI(m Meter) Inch {
	return Inch(LMtoC(m) / OneF_CM)
}

//LMtoF Convert to
func LMtoF(m Meter) Feet {
	return Feet(LMtoC(m) / OneF_CM)
}
