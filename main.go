package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Fruits": 500,
		"Tea":    200,
	}
	menu["Desert"] = 300
	menu["Салат"] = 600

	delete(menu, "Салат")

	value, ok := menu["Салат"]
	fmt.Println("Салат", value, ok)

	for item, price := range menu {
		fmt.Printf("%s стоит %.2f рублей\n", item, price)
	}
}
