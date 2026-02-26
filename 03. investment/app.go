package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 5
	expectedReturnRate := 5.5

	// In this code, investmentAmount is set as float64 because it represents a monetary value, which can have decimal places. Using float64 allows for more precision when dealing with financial calculations.
	// On the other hand, years is an integer value that represents the number of years for the investment. However, when calculating the future value of the investment, we need to perform a multiplication operation that involves both the expected return rate and the number of years. To ensure that the calculation is accurate and does not lose precision, we cast years to float64 before performing the multiplication. This allows us to maintain the precision of the calculation and avoid any potential issues with integer division or rounding errors.
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
