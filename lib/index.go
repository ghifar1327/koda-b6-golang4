package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func KonversiApp(){

	scanner := bufio.NewScanner(os.Stdin)


	defer func() {
		if r := recover(); r != nil {
			if r == "400" {
				fmt.Println(r,": Input tidak valid")
			} else if r == "404" {
				fmt.Println(r,": Menu tidak tersedia")
			} else {
				fmt.Println("Terjadi kesalahan:", r)
			}

			fmt.Println("1. kembali ke input")
			fmt.Println("2. akhiri program")
			fmt.Print("Masukan pilihan: ")

			scanner.Scan()
			choice := scanner.Text()

			if choice == "2" {
				fmt.Println("Program selesai")
				os.Exit(0)
			}
		}
	}() 

	for {

		fmt.Print("\nMasukan input celcius: ")
		scanner.Scan()
		text := scanner.Text()

		i, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic("400")
		}

		fmt.Println("Nilai celcius =", i)
		Menu()

		fmt.Print("Pilih target konversi (1/2/3): ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			Fahrenhit(i)
		case "2":
			Reamur(i)
		case "3":
			Kelvin(i)
		case "4":
			fmt.Println("keluar dari program")
		default:
			panic("404")
		}
	}
}

