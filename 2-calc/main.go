package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var operation string
var finalNumbers []float64

func main() {
	for {
		operation, err := enterOperation()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Вы выбрали операцию:", operation)
		break
	}

	for {
		finalNumbers, err := enterNumbers()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Вы ввели следующие числа:", finalNumbers)
		break
	}

	result := calculation(operation, finalNumbers)

	switch operation {
	case "SUM":
		fmt.Println("Сумма чисел равна:", result)
	case "AVG":
		fmt.Println("Среднее чисел равно:", result)
	case "MED":
		fmt.Println("Медиана чисел равна:", result)
	}
}

func enterOperation() (string, error) {
	fmt.Println("Введите операцию из предложенных (AVG, SUM, MED):")
	_, _ = fmt.Scan(&operation)
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		return "", errors.New("операция должна совпадать с одной из предложенных")
	}
	return operation, nil
}

func enterNumbers() ([]float64, error) {
	var listOfNumbers string
	var numbers []string
	fmt.Println("Введите числа через запятую")
	_, _ = fmt.Scan(&listOfNumbers)
	numbers = strings.Split(listOfNumbers, ",") // срез из строк
	for _, value := range numbers {
		element, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Error!")
		}
		finalNumbers = append(finalNumbers, element)
	}
	return finalNumbers, nil
}

func calculation(operation string, numbers []float64) float64 {
	var result float64
	var preResult float64
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	switch operation {
	case "AVG":
		for _, value := range numbers {
			preResult = preResult + value
		}
		result = preResult / float64(len(numbers))
	case "SUM":
		for _, value := range numbers {
			result = result + value
		}
	case "MED":
		fmt.Println("Упорядоченный ряд:", numbers)
		if len(numbers)%2 != 0 {
			result = numbers[int(math.Round(float64(len(numbers))/2))-1]
		} else if len(numbers)%2 == 0 {
			result = (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
		}
	}
	return result
}
