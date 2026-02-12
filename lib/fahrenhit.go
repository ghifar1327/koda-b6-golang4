package lib

import "fmt"

func Fahrenhit(c float64) {
	f := (c * 9 / 5) + 32
	fmt.Printf("%.2f celcius = %.2f fahrenheit\n", c, f)
}