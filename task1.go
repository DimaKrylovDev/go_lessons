package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run . <температура в °C>")
		fmt.Println("Пример: go run . 36.6")
		os.Exit(1)
	}

	celsiusStr := os.Args[1]

	celsius, err := strconv.ParseFloat(celsiusStr, 64)
	if err != nil {
		fmt.Printf("Ошибка: '%s' не является корректным числом\n", celsiusStr)
		os.Exit(1)
	}
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f °F\n", fahrenheit)
}