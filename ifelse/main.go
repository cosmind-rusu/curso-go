package main

import "fmt"

func main() {
	// CREAMOS LA VARIABLE
	var n int

	// ENTRADA DE DATOS
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&n)

	// PORCESOS Y SALIDA DE DATOS

	if n == 0 {
		fmt.Println(n, "es neutro")
	} else if n % 2 == 0 {
		fmt.Println(n, "es par")
	} else {
		fmt.Println(n, "es impar")
	}

}