package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___Калькулятор индекса массы тела___")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println("Введены некорректные параметры для расчета")
			fmt.Println("Приложение завершило работу")
			break
			//panic("Не заданы параметры для расчета")
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		} else {
			continue
		}
	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела составляет: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case imt >= 16 && imt < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case imt >= 18.5 && imt < 25:
		fmt.Println("У Вас нормальный вес")
	case imt >= 25 && imt < 30:
		fmt.Println("У Вас избыточный вес")
	default:
		fmt.Println("У Вас ожирение")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Вы хотите сделать еще один расчет (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	} else {
		return false
	}
}
