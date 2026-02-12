package lib

import "fmt"

func Kelvin(c float64) {
	k := c + 273.15
	fmt.Printf("%.2f celcius = %.2f kelvin\n", c, k)
}