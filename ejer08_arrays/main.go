package main

import "fmt"

var tabla [10]int

func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	tabla2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	// matriz
	var matriz [5][7]int
	matriz[3][5] = 5
	fmt.Println(matriz)

	// Slices
	// Son vectores dinamicos, puedo cambiar las dimenciones en tiempo de ejecución
	// Si no le coloco ningun tipo de longitud es un slice
	slice2 := []int{2, 4, 5}
	fmt.Println(slice2)
	variante2()
	variante3()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	// desde la posicion 3 hasta la ultima cree un slice
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
