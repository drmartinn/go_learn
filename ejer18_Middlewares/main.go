package main

import "fmt"

/* Un midlleware es un interceptor que permiten ejecutar intrucciones
comunes a varias funciones que reciben y devuelven los mismos tipos de variables */
var result int

func main() {
	fmt.Println("Inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Printf("Inicio de operacion con valor a %d y valor b %d\n", a, b)
		return f(a, b)
	}
}
