package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func main() {
	// Создание объекта структуры
	car := Car{
		Brand: "Porsche",
		Year:  2020,
	}
	updateCarYear(car)

	fmt.Println(car)
}

func updateCarYear(car Car) {
	car.Year++
	fmt.Println(car)

}
