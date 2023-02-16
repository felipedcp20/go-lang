package main

import "fmt"

func main() {
	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	println("\n")

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	println("\n")
	println("\n")

	//for while
	counter := 0
	for counter < 10 {
		println(counter)
		counter++
	}

	//For Forever
	counterforever := 0
	for {
		println(counterforever)
		counterforever++
	}

	//For RANGE
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

}
