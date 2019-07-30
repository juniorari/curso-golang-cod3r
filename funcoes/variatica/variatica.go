package main

import "fmt"

func media(numeros ...int) float64 {
	total := 0
	for _, num := range numeros {
		total += num
	}
	println(numeros, len(numeros))
	return float64(total / len(numeros))
}

func main() {
	fmt.Printf("Média: %f", media(4, 2, 3, 4))
}
