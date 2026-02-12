package lib

import "fmt"

func Reamur(c float64) {
	r := c * 4 / 5
	fmt.Printf("%.2f celcius = %.2f reamur\n", c, r)
}