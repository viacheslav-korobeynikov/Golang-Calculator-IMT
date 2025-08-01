package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	var userHeight float64
	var userKg float64
	fmt.Println("___Калькулятор индекса массы тела___")
	fmt.Println("Введите свой рост (в метрах): ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Println("Ваш индекс массы тела составляет: ")
	fmt.Print(IMT)
}
