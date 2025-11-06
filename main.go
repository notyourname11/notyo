package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal Animal
}

func (d Dog) Bark() {
	fmt.Printf("%s гавкает\n", d.Animal)
}

func (a Animal) Speak() {
	fmt.Printf("%s издает какой-то звук\n", a.Name)
}

func main() {
	dog := Dog{
		Animal: Animal{
			Name: "Rex",
		},
	}

	dog.Bark()
	dog.Animal.Speak()

}
