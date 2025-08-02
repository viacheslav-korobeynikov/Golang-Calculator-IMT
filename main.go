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

	switch {
	case IMT < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case IMT >= 16 && IMT < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case IMT >= 18.5 && IMT < 25:
		fmt.Println("У Вас нормальный вес")
	case IMT >= 25 && IMT < 30:
		fmt.Println("У Вас избыточный вес")
	default:
		fmt.Println("У Вас ожирение")
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела составляет: %.0f", imt)
	fmt.Println(result)
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
