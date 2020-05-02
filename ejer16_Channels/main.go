package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")
	// Esperando a que canal1 obtenga un valor
	msg := <-canal1
	fmt.Printf("Canal 1 %v", msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	//Asignando un valor al canal1, la diferencia de valor entre inicio y final
	canal1 <- final.Sub(inicio)
}

// Un chanel es un espacio de memoria de dialogo, debe ser un tipo de dato, transporta una variable
