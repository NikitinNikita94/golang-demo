package main

import (
	"fmt"
	"math"
)

const IMT_POW = 2

func main() {
	fmt.Println("Привет, введите ваш рост в сантиметрах и вес")

	userHeight, userWeight, err := readUserParam()
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	imt := calculateImt(userHeight, userWeight)

	fmt.Printf("Ваш индекс массы тела составляет %.2f\n", imt)
	fmt.Printf("Длина строки Golang %d", len("Golang"))
}

func calculateImt(height float64, weight float64) float64 {
	return weight / math.Pow(height/100, IMT_POW)
}

func readUserParam() (height float64, weight float64, err error) {
	fmt.Print("Рост: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при вводе роста: %v", err)
	}

	fmt.Print("Вес: ")
	_, err = fmt.Scan(&weight)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при вводе веса: %v", err)
	}

	return height, weight, nil
}
