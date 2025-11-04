package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("нельзя делить на ноль")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age == 0 {
		return errors.New("некорректный возраст")
	}
	return nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil { // Если ошибка не nil, то есть если есть ошибка
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
	err = validateAge(0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Возраст корректен")

	}
}
