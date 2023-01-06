package main

import "fmt"

func main() {
	fmt.Println("OPERADORES DE ASIGNACION")
 	a := 5
	//a = a + 3 // Los operadores en asignacion simplifican las operaciones ANTES
	a += 3 // Los operadores en asignacion simplifican las operaciones DESPUES
	fmt.Println(a)
	fmt.Println("OPERADORES DE ASIGNACION")

	fmt.Println("\n")

	fmt.Println("OPERADORES DE INCREMENTO")
	//operadores de incremento
	b := 1
	b++
	fmt.Println(b)
	for b := 0; b < 4; b++ { // ++ es una instrucción.
        fmt.Println(b)
    }
	fmt.Println("OPERADORES DE INCREMENTO")

	



}