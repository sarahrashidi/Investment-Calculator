package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 6.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Formated Future Value & Real future value
	formatedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future value(adjusted for Inflation): %.1f", futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
	


}


