package main

import "fmt"

// funcion anonima
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", calculo(5, 7))
	// Se puede re definir la funcion calculo
	// se debe respetar el mismo tipo y numero de datos de entrada
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 2 = %d \n", calculo(6, 2))
	calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividor 6 / 2 = %d \n", calculo(6, 2))

	operaciones()

	/* CLOSURES */
	tablaDel := 2
	MiTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}

}

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// Closures
// Funciones anonimas, proteccion del codigo
// pueden acceder a variables por fuera de la funcion
// Una funcion que devuelve otra funcion que devuelve un entero
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
