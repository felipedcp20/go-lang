package main

import "fmt"

func main()  {
	
	switch modulo := 3 % 2; modulo {
	case 0:
		fmt.Println("el numero es par")
	default:
		fmt.Println("el numero es impar")
	}


	// SWITCH SIN CONDICION
	value := 100
	switch {
	case value > 100:
		fmt.Println("es mayor que 100")
	case value < 100:
		fmt.Println("el valor es menor que 100")
		default:
			fmt.Println("es valor es 100")	
	}
}



