package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese un número secreto")
		fmt.Scanf("%d", &numero)
		if numero == 100 {
			// Para salir del ciclo
			break
		}
	}

	var index = 0
	var index2 = 50
	for index < 10 {
		fmt.Printf("\n valor de index: %d y otra variable iniciada en %d", index, index2)
		if index == 5 {
			fmt.Printf(" multplicamo por 2 \n")
			index = index * 2
			// el continue envia el contro al inicio del cilo for, es decir no imprime pasó por aquí
			continue
		}
		fmt.Printf("    Pasó por aquí \n")
		index++
	}

	var l int
	// Esto es una etiqueta
RUTINA:
	for l < 10 {
		if l == 4 {
			l = l + 2
			fmt.Println("Voy a RUTINA sumando 2 a l")
			goto RUTINA
		}
		fmt.Printf("valor de l: %d \n ", l)
		l++
	}

}
