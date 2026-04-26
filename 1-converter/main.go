package main

import "fmt"

const USDtoEUR = 0.8532
const USDtoRUB = 75.53

func main() {
	EURtoRUB := EurToRubRate(USDtoEUR, USDtoRUB)
	fmt.Println("Курс Евро к Рублю: ", EURtoRUB)
	UserInput()
}

func EurToRubRate(USDtoEUR float64, USDtoRUB float64) float64 {
	EURtoRUB := USDtoRUB / USDtoEUR
	return EURtoRUB
}

func UserInput() (float64, string, string) {
	var amount float64
	var inCurrency string
	var outCurrency string
	fmt.Println("Введите сумму денег: ")
	fmt.Scan(&amount)
	fmt.Println("Введите исходную валюту (RUB, USD, EUR): ")
	fmt.Scan(&inCurrency)
	fmt.Println("Введите целевую валюту (RUB, USD, EUR): ")
	fmt.Scan(&outCurrency)

	fmt.Println("Введенные данные:")
	fmt.Println("Сумма денег:", amount, inCurrency)
	fmt.Println("Исходная валюта:", inCurrency)
	fmt.Println("Целевая валюта:", outCurrency)

	return amount, inCurrency, outCurrency
}

func CalculateMoney(amount float64, inCurrency string, outCurrency string) float64 {

}
