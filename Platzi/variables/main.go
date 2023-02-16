package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Declaracion de constantes

	const pi float64 = 3.14
	var pi2 = 3.14
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaraciòn de variables Enteras

	base := 12 // se declara una variable con un valor pero sin su tipo de valor
	var altura int = 14
	var area int
	area = base * altura // así se le asigna un valor a una variable previamente declarada
	fmt.Println("la altura es: ", area)

	// Zero Value
	// En otros lenguajes al no declar el valor de una variable toma por defecto NULL en Go toma los siguientes:
	var a int     // 0
	var b string  // vacio
	var c bool    // false
	var d float64 // 0
	fmt.Println(a, b, c, d)

	//Area Cuadrado
	const baseCuadrado int = 10
	AreaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("el area del cuadrado es: ", AreaCuadrado)

	//OPERADORES ARITMETICOS

	x := 3
	y := 2
	//SUMA
	result := x + y
	fmt.Println("el reseultado de la suma es: ", result)

	//RESTA
	result = x - y
	fmt.Println("el reseultado de la resta es: ", result)

	//MULTIPLICACION
	result = x * y
	fmt.Println("el reseultado de la multiplicacion es: ", result)

	//DIVISION

	result2 := float64(x) / float64(y)
	fmt.Println("el reseultado dTe la division es: ", result2)

	//MODULO
	result = x % y
	fmt.Println("el reseultado del modulo es: ", result)

	//INCREMENTAL
	x++
	fmt.Println("INCREMENTAL", x)

	//DECREMENTAL
	x--
	fmt.Println("DECREMENTAL", x)

	//PACKAGE FMT
	nombre := "platzi"
	cursos := 500

	fmt.Println("hello")                                      // REALIZA SALTO DE LINEA
	fmt.Printf("%v tiene mas de %d cursos\n", nombre, cursos) // para concatenar textos dentro de un PRINT
	// \n al final para realizar salto de linea
	// para concatenar textos dentro de un PRINT
	// %s = es un string
	// %d = es un entero
	// %v = si no sabes cual es el valor
	fmt.Printf("%T %T \n", nombre, cursos)
	// %T = imprime el tipo de variable (int,string,boolean)

	//concatenar variables en un string y almacenarlos en una variable
	message := fmt.Sprintf("%v tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	message2 := fmt.Sprintf("%T %T %T %T\n", a, b, c, d)
	fmt.Println(message2)
}
