package main

import (
	"errors"
	"fmt"
)

const USDtoEUR = 0.8532
const USDtoRUB = 75.53
const EURtoRUB = 88.53

var inCurrency string
var outCurrency string
var amount float64

func main() {
	fmt.Println("Вас приветствует конвертер валют!")
	for {
		inCurrency, err := EnterInCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Исходная валюта: ", inCurrency)
		break
	}
	for {
		outCurrency, err := EnterOutCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if outCurrency == inCurrency {
			fmt.Println("Целевая валюта не должна совпадать с исходной")
			continue
		}
		fmt.Println("Целевая валюта: ", outCurrency)
		break
	}
	for {
		amount, err := EnterAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Сумма денег: ", amount, inCurrency)
		break
	}
	result := CalculateMoney(amount, inCurrency, outCurrency)
	fmt.Println("Результат: ", result, outCurrency)
}

func EnterInCurrency() (string, error) {
	fmt.Println("Введите исходную валюту (RUB, USD, EUR): ")
	_, _ = fmt.Scan(&inCurrency)
	if inCurrency != "RUB" && inCurrency != "USD" && inCurrency != "EUR" {
		return "", errors.New("валюта должна совпадать с одной из предложенных\n")
	}
	return inCurrency, nil
}

func EnterOutCurrency() (string, error) {
	switch {
	case inCurrency == "RUB":
		fmt.Println("Введите целевую валюту (USD, EUR): ")
	case inCurrency == "USD":
		fmt.Println("Введите целевую валюту (RUB, EUR): ")
	case inCurrency == "EUR":
		fmt.Println("Введите целевую валюту (RUB, USD): ")
	}
	_, _ = fmt.Scan(&outCurrency)
	if outCurrency != "RUB" && outCurrency != "USD" && outCurrency != "EUR" {
		return "", errors.New("валюта должна совпадать с одной из предложенных\n")
	}
	return outCurrency, nil
}

func EnterAmount() (float64, error) {
	fmt.Println("Введите сумму денег (положительное число, используя цифры): ")
	_, err := fmt.Scanf("%f", &amount)
	if err != nil {
		return 0, errors.New("сумма должна быть введена цифрами\n")
	} else if amount < 0 {
		return 0, errors.New("сумма должна быть положительной\n")
	}
	return amount, nil
}

func CalculateMoney(amount float64, inCurrency string, outCurrency string) float64 {
	var money float64
	switch {
	case inCurrency == "RUB" && outCurrency == "EUR":
		money = amount / EURtoRUB
	case inCurrency == "RUB" && outCurrency == "USD":
		money = amount / USDtoRUB
	case inCurrency == "USD" && outCurrency == "EUR":
		money = amount * USDtoEUR
	case inCurrency == "USD" && outCurrency == "RUB":
		money = amount * USDtoRUB
	case inCurrency == "EUR" && outCurrency == "RUB":
		money = amount * EURtoRUB
	case inCurrency == "EUR" && outCurrency == "USD":
		money = amount / USDtoEUR
	}
	return money
}
