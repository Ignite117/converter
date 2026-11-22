package main

import (
	"fmt"
	"strconv"
)

const USD_TO_EUR = 0.91
const USD_TO_RUB = 82.83
const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB

func getUserInputStValut() string {
	var per string
	fmt.Scan(&per)
	return per
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
		stValut = getUserInputStValut()
		if stValut != "USD" && stValut != "RUB" && stValut != "EUR" {
			fmt.Println("Неверная валюта")
			continue
		}
		fmt.Println("Сумма денег:")
		count, _ = strconv.ParseFloat(getUserInputStValut(), 64)
		if count < 0 {
			fmt.Println("Некоректная сумма")
			continue
		}
		fmt.Println("Итоговая валюта(USD/RUB/EUR):")
		endValut = getUserInputStValut()
		if endValut != "USD" && endValut != "RUB" && endValut != "EUR" {
			fmt.Println("Неверная валюта")
			continue
		}
		if stValut == endValut {
			fmt.Println("Одинаковая валюта")
			continue
		}
		total := calculate(count, stValut, endValut)
		fmt.Println(total)

	}
}
