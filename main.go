package main

import "fmt"

type Movable interface {
	Move() string
}

func (c Car) Move() string {
	return c.Brand + " Мчится со скоростью " + c.Speed
}

type Car struct {
	Brand string
	Speed string
}

type Person struct {
	Name string
	Age  string

	Speed string
}

func (p Person) Move() string {
	return p.Name + " Идет со скоростью " + p.Speed
}

func printMovement(m Movable) {
	fmt.Println(m.Move())
}

func main() {

	Porsche := Car{Brand: "Porsche", Speed: "250"}
	Andrey := Person{Name: "Andrey", Age: "20", Speed: "10"}
	printMovement(Porsche)
	printMovement(Andrey)

}
