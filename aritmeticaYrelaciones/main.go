package main

import "fmt"

func main() {
	// #################################### OPERADORES ARITMETICOS ####################################

	fmt.Println("OPERADORES ARITMETICOS")

	// CREAMOS VARIABLES
	var a, b int
	// ENTRADA DE DATOS
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&b)

	// PROCESOS
	suma := a + b
	resta := a - b
	multiplicacion := a * b
	division := a / b
	resto := a % b

	// SALIDA DE DATOS
	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)	
	fmt.Println("La multiplicacion es: ", multiplicacion)
	fmt.Println("La division es: ", division)
	fmt.Println("El resto es: ", resto)

// #################################### OPERADORES ARITMETICOS ####################################

// #################################### OPERADORES RELACIONALES ####################################

	// CREAMOS VARIABLES
	var c, d int
	// ENTRADA DE DATOS
	fmt.Println("Ingresa c: ")
	fmt.Scanln(&c)
	fmt.Print("Ingrese d: ")
	fmt.Scanln(&d)

	// PROCESOS
	igualdad := c == d
	distintos := c != d
	mayorQue := c > d
	menorQue := c < d
	mayorIgual := c >= d
	menorIgual := c <= d

	// SALIDA DE DATOS
	fmt.Println("Son iguales? ", igualdad)
	fmt.Println("Son distintos? ", distintos)
	fmt.Println("c es mayor que d? ", mayorQue)
	fmt.Println("c es menor que d? ", menorQue)
	fmt.Println("c es mayor o igual que d? ", mayorIgual)
	fmt.Println("c es menor o igual que d? ", menorIgual)

}