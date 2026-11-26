package main

import (
	"fmt"
	"strconv"
)

const USD_TO_EUR = 0.91
const USD_TO_RUB = 82.83
const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB

func getUserInputValut(prev string) string {
	var per string
	for {
		fmt.Scan(&per)
		if prev == per {
			fmt.Println("Одинаковая валюта")
			continue
		}
		if ParseValut(per) {
			break
		}
	}
	return per
}
func getUserInputSum() float64 {
	var sum string
	var sumf float64
	var err error
	for {
		fmt.Scan(&sum)
		sumf, err = strconv.ParseFloat(sum, 64)
		if err == nil && sumf > 0 {
			break
		} else {
			fmt.Println("Невалидная сумма")
		}
	}
	return sumf
}
func ParseValut(Valut string) bool {
	if Valut != "USD" && Valut != "RUB" && Valut != "EUR" {
		fmt.Println("Неверная валюта")
		return false
	}
	return true
}
func calculate(count float64, stValut string, endValut string) float64 {
	switch {
	case stValut == "USD" && endValut == "EUR":
		return USD_TO_EUR * count
	case stValut == "EUR" && endValut == "USD":
		return 1 / USD_TO_EUR * count
	case stValut == "USD" && endValut == "RUB":
		return USD_TO_RUB * count
	case stValut == "RUB" && endValut == "USD":
		return 1 / USD_TO_RUB * count
	case stValut == "EUR" && endValut == "RUB":
		return EUR_TO_RUB * count
	case stValut == "RUB" && endValut == "EUR":
		return 1 / EUR_TO_RUB * count
	}
	return 0.0
}
func main() {
	var count float64
	var stValut, endValut string
	for {
		fmt.Println("Исходная валюта(USD/RUB/EUR):")
		stValut = getUserInputValut("")
		fmt.Println("Сумма денег:")
		count = getUserInputSum()
		fmt.Println("Итоговая валюта(USD/RUB/EUR):")
		endValut = getUserInputValut(stValut)
		fmt.Println("Итог:")
		total := calculate(count, stValut, endValut)
		fmt.Println(total)

	}
}
