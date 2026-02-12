package lib

import "fmt"

func Menu() {
	list := []string{
		"fahreinhit",
		"reamur",
		"kelvin",
	}
	for i, v := range list {
		fmt.Println(i+1, v)
	}
}

