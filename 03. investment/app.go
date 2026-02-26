package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedResturnRate = 5.5
	var years float64 = 5

	// In this code, investmentAmount is set as float64 because it represents a monetary value, which can have decimal places. Using float64 allows for more precision when dealing with financial calculations.
	// On the other hand, years is an integer value that represents the number of years for the investment. However, when calculating the future value of the investment, we need to perform a multiplication operation that involves both the expected return rate and the number of years. To ensure that the calculation is accurate and does not lose precision, we cast years to float64 before performing the multiplication. This allows us to maintain the precision of the calculation and avoid any potential issues with integer division or rounding errors.
	var futureValue = investmentAmount * math.Pow(1+expectedResturnRate/100, years)
	fmt.Println(futureValue)
}
