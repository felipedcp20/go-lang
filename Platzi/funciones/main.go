package main

import "fmt"

func saludarPersonas(nombre string) {
	fmt.Printf("hola %s espero que estes bien\n", nombre)
}

func multiArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func aumentDoble(numero int) int {
	value := numero * 2
	return value
}

func dobleReturn(a int) (b,c int){
	return a, a *2
}

func main() {
	saludarPersonas("felipe")
	saludarPersonas("Daniela")
	saludarPersonas("Dilan")
	multiArgument(1, 3, "daniela")

	fmt.Println(aumentDoble(9))
	value1, _ := dobleReturn(1)
	fmt.Println(value1)

}
