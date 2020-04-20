package main

import "fmt"

func main() {
	paises := make(map[string]string)
	fmt.Println(paises)
	// key / value
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos aires"
	fmt.Println(paises)

	//Reservando espacios , 5
	paises2 := make(map[string]string, 5)
	fmt.Println(paises2)
	campeonato := map[string]int{
		"Barcelona":   39,
		"Real madrid": 42,
		"Chivas":      37,
		"Boca junior": 30}
	// Agregando un item
	campeonato["River plate"] = 25
	// Actualizando un valor
	campeonato["Chivas"] = 99
	// Eliminando un elemento
	delete(campeonato, "Real madrid")
	fmt.Println(campeonato)

	// Recorriendo un mapa
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}
	// Validar si existe un key en un map
	puntaje, existe := campeonato["Mineiro"]
	// %t es para los booleanos
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t\n", puntaje, existe)
}
