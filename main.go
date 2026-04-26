package main

import "fmt"

const USDtoEUR = 0.8532
const USDtoRUB = 75.53

func main() {
	EURtoRUB := EurToRubRate(USDtoEUR, USDtoRUB)
	fmt.Print("Курс Евро к Рублю: ", EURtoRUB)
}

func EurToRubRate(USDtoEUR float64, USDtoRUB float64) float64 {
	EURtoRUB := USDtoRUB / USDtoEUR
	return EURtoRUB
}
