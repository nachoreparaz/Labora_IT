package main

import (
	"fmt"
	"math"
)

func main() {
	// esfera()
	// raiz()
}

// Define una función que calcule el volumen de una esfera.
func esfera() {
	radio := 3
	volumen := (4 / 3) * math.Pi * math.Pow(3, float64(radio))
	fmt.Println("El volumen de la esfera es:", volumen)
}

func raiz() {
	number := 81
	raiz := math.Sqrt(float64(number))
	fmt.Println("La raiz cuadrada de", number, "es:", raiz)
}
