package main

import "fmt"

func main() {
	menu := `Elija una operación
	[1] suma
	[2] resta
	[3] multiplicación
	[4] división
	`

	fmt.Println(menu)
	var eleccion int
	fmt.Scanln(&eleccion)

	switch eleccion {
	case 1:
		fmt.Println(sum())

	case 2:
		fmt.Println(rest())

	case 3:
		fmt.Println(mult())

	case 4:
		fmt.Println(div())
	}

}
