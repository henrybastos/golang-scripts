package main

import (
	"fmt"
)

func main () {
	var x [4]string
	x[0] = "Henry"
	x[1] = "Bastos"
	x[2] = "da"
	x[3] = "Silva"
	fmt.Println(x[0])
	fmt.Println(len(x))
}