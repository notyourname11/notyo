package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func updateCarYearPointer(car *Car) {
	car.Year++

}

func main() {
	car := Car{Brand: "Porsche", Year: 2020}
	updateCarYearPointer(&car)
	fmt.Println(car)
}
