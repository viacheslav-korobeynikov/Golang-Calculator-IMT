package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	userHeight := 1.8
	var userKg float64 = 100
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
