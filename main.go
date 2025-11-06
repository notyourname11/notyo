package main

import "fmt"

func main() {
	in := []int{5, 4, 9, 2, 7, 6}
	for i := len(in) - 1; i >= 0; i-- {
		if in[i]%2 != 0 {
			in = append(in[:i], in[i+1:]...)
		}
	}
	fmt.Println((in))

}
