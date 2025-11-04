package main

import "fmt"

func main() {
	type Laptop struct {
		Brand string
		Price int
	}

	laptop := []Laptop{
		{Brand: "Apple", Price: 300}, {Brand: "HP", Price: 200},
	}
	laptop[1] = Laptop{Brand: "HP", Price: 180}
	newLaptop := Laptop{Brand: "MSI", Price: 200}
	laptop = append(laptop, newLaptop)

	for index, laptop := range laptop {
		fmt.Printf("%d. %s by %d\n", index, laptop.Brand, laptop.Price)
	}
}
