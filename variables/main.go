package main

import "fmt"

// LAS CONSTANTES SE PUEDEN DECLARAR FUERA DE LA FUNCION
const nombre string = "Cosmin :)"

func main() {
	// PARA DEFINIR UNA VARIABLE USAREMOS LA PALABRA "var"
	 // Definimos una variable llamada nombre y le asignamos un valor
	var nombre string = "Cosmin"
	//ahora la llamamos con el metodo fmt y su operador Println
	fmt.Println(nombre)
	// ahora comprobamos que la variable cambia su valor si mas abajo del codigo cambia el contenido
	nombre = "Eva"
	//ahora la llamamos con el metodo fmt y su operador Println
	fmt.Println(nombre)
	// veremos que ahora el nombre que nos saldra por consola sera "Eva"

	

	// Ahora definimos varias variables al mismo tiempo
	var a, b, c string = "Manuel", "Manuel", "Alex"
	var d, e int = 3, 22
	//ahora la llamamos con el metodo fmt y su operador Println
	fmt.Println(a, b, c, d, e)

	// lo que hemos hecho ahora es definir un objeto de variables y mas abajo le hemos dado valor a algunas de ellas
	var (
		pi float64
		booleano bool
		cadena = "Texto de la cadena"
		edad = 18
	)

	pi = 3.1415
	booleano = false

	//ahora la llamamos con el metodo fmt y su operador Println
	fmt.Println(pi, booleano, cadena, edad)

	// AQUI DEFINIMOS UNA VARIABLE SIN USAR LA INSTRUCION "var" usando la carida de GO :=
	// Los 2 puntos que ponemos al principio los colocamos para la definicion de la variable pero luego ya no hara falta, ya que ya esta definida y solo usaremos el nombre y el =
	v1 := 23
	v1 = 57

	v2 := "Texto de variable simplificada"

	fmt.Println(v1,v2)

	// AHORA TRABAJAREMOS CON CONSTANTES
	// las constantes no cambian como las variables, y las definimos con la palabra reservada "const"
	
	const n = 100
	
	//ahora la llamamos con el metodo fmt y su operador Println
	fmt.Println(n, nombre)

}