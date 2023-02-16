package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var numero1 int = 1
	var numero2 int = 3


	//if
	if numero1 == numero2 {
		fmt.Println("son iguales")
	} else {
		fmt.Println("son diferentes")
	}

	//if with and

	if numero1 == 1  && numero2 == 3{
		fmt.Println("los numeros son 1 y 3")
	}else {
		println("los numeros no son 1 y 3")
	}

	//if with or

	if numero1 == 1 || numero2 ==3{
		fmt.Println("alguno de los numeros es 1 o 3")
	}else {
		println("ninguno de los numeros es 1 o 3")
	}

	//convertir texto a numero

	value, err := strconv.Atoi("1233")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("valor: ", value)

}
