package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.91
	const USD_TO_RUB = 82.83
	const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB
	fmt.Printf("EUR to RUB: %v", 1/USD_TO_EUR*USD_TO_RUB)
}
