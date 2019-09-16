package main

import (
	"fmt"

	"github.com/juniorari/goarea"
)

// Necessário rodar o comando:
/*
$ go get -u github.com/juniorari/goarea
*/
func main() {
	fmt.Println(goarea.Circ(6.0))
	fmt.Println(goarea.Rect(5.0, 2.0))
	// fmt.Println(area._TrianguloEq(5.0, 2.0))
}
