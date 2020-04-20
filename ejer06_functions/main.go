package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(3)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(tres(5))
	fmt.Println("***********")
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5, 46, 1))
	fmt.Println(calculo(1))
	fmt.Println(calculo(1, 2, 3, 5))

}

// funcion que retorna un unico valor
func uno(numero int) int {
	return numero * 2
}

// funcion que retorna dos valores
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// funcion que retorna una variable
func tres(numero int) (z int) {
	z = numero * 2
	return
}

// Parametros variables, no se sabe cuales son los parametros de entrada
// Cuando comienza con Mayus es publica, de lo contrario es privada
func calculo(numero ...int) int {
	total := 0
	// No se puede obviar parametros
	// el _ sirve para indicar que ese parametro no lo voy a utilizar
	// No reserva memoria en el _
	for _, num := range numero {
		total = total + num
	}
	return total
}
