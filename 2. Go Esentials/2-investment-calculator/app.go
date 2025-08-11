package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentValue = 1000
	var returnRate = 5.5
	var years = 10

	var futureValue = float64(investmentValue) * math.Pow(1 + returnRate / 100, float64(years))
	fmt.Printf("Future value of investment: $%.2f\n", futureValue)
}