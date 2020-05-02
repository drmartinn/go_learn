package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, error := ioutil.ReadFile("./archivo.txt")
	if error != nil {
		fmt.Println("Hubo un error leyendo el archivo")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, error := os.Open("./archivo.txt")
	if error != nil {
		fmt.Println("Hubo un error leyendo el archivo")
	} else {
		scanner := bufio.NewScanner(archivo)
		//Se reliaza la lectura de la linea
		for scanner.Scan() {
			//Se obtiene la linea actual convertirla en texto
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + " \n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\nEsta es una segunda linea") == false {
		fmt.Printf("error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
