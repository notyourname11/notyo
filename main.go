package main

import "fmt"

// SumInts возвращает сумму всех чисел nums.
// Если срез пуст, должна вернуться 0.
func SumInts(nums []int) int {
	SumInts := 0
	for index, _ := range nums {
		SumInts += nums[index]
	}
	return SumInts

}

func main() {
	fmt.Println(SumInts([]int{1, 2, 3}))
	fmt.Println(SumInts([]int{}))
	// Ожидаемый вывод:
	// 6
	// 0
}
