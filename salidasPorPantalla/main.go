package main

import "fmt"

func main () {
	// ESTE PARAMETRO NO HACE SALTO DE LINEA
	// PARA PODER HACER UN SALTO DE LINEA EN UNA SALIDA POR PANTALLA PONDRIAMOS \t para tabular y \n para salto de linea
	fmt.Print("Hola\t mundo \n")

	// DECLARAMOS VARIABLES
	nombre := "Cosmin"
	edad := 18
	pi := 3.1415

	// LAS MOSTRAMOS
	fmt.Println("Nombre: ", nombre, "Edad: ", edad, "\nPi: ", pi)

	// USAMOS EL FORMATEO DE GO PARA PODER CONCATENAR DE MANERA SENCILLAS LAS VARIABLES DECLARADAS

/* 	%s para mostrar datos de tipo string
	%d para mostrar datos de tipo entero
	%f paÍa mostrar datos de tipo flotante */

	fmt.Printf("Nombre: %s Edad: %d \nPi: %f", nombre, edad, pi)
}