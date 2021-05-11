package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 25, 34, 35, 51, 52, 53, 54}
	objetivo := 52
	size := len(list)
	result := binarySearch(list, 0, size-1, objetivo)
	if result == -1 {
		fmt.Println("El elemento no se ha encontrado")
	} else {
		fmt.Println("El elemento se ha encontrado en la posicion", result)
	}
}

func binarySearch(list []int, left int, rigth int, objetive int) int {
	fmt.Println(left)
	fmt.Println(rigth)
	if rigth >= left {
		mid := left + (rigth-left)/2
		if list[mid] == objetive {
			return mid
		}
		if list[mid] > objetive {
			return binarySearch(list, left, mid-1, objetive)
		} else {
			return binarySearch(list, mid+1, rigth, objetive)
		}
	}
	return -1
}
