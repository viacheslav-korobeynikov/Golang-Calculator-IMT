package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___Калькулятор индекса массы тела___")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела составляет: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Println("Введите свой рост (в сантиметрах): ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)
	return userKg, userHeight

}
