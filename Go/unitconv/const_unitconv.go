// Unit converter program

package unitconv

import "fmt"

// temperature
// Celsius
// Fahrenheit

//length
// feet
// meters
// cm
// inch

//weight
//pounds
//kilograms

type Common float64

type Kelvin float64
type Celsius float64
type Fahrenheit float64

type KG float64
type LBS float64

type Inch float64
type CM float64
type Feet float64
type Meter float64

const (
	ZeroK         = -273.15
	AbsoluteZeroC = -273.15
)

const (
	OneIN_CM = 2.54
	OneF_CM  = 30.48
	OneM_CM  = 100
)

// Common c = 5 in Celsius
// string s = °C
// concat will return 5°C
func (c Common) concat(s string) string {
	str := new(string)
	*str = fmt.Sprintf("%v%s", c, s)
	return *str
}
