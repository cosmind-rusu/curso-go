package main

import "fmt"

func main() {
	// CREAMOS LAS VARIABLES
	var (
		a int
		salida string
	)
		
	// ENTRADA DE DATOS

	fmt.Print("Ingrese un numero del 1 al 5: ")
	fmt.Scanln(&a)

	// PROCESOS

	switch a {
	case 1:
		salida = "UNO"
	case 2:
		salida = "DOS"
	case 3:
		salida = "TRES"
	case 4:
		salida = "CUATRO"
	case 5:
		salida = "CINCO"

	default:
		salida = "NO ESTA ENTRE EL 1 Y EL 5"
	}

	// SALIDA DE DATOS

	fmt.Println(a, salida)
}
