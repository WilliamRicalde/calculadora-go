package main

import "fmt"

func sum() int {
	var a int
	var b int

	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&b)

	return a + b
}

func rest() int {
	var a int
	var b int

	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&b)

	return a - b
}

func mult() int {
	var a int
	var b int

	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&b)

	return a * b
}

func div() int {
	var a int
	var b int

	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&b)

	return a / b
}
