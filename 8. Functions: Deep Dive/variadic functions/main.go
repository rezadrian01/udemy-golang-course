package main

import "fmt"

func main(){
	numbers := []int{2, 3, 1}

	sum := sumUp(1, 10, 15, 43)
	anotherSum := sumUp(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumUp(numbers ...int) int{
	sum := 0
	for _, val := range numbers{
		sum += val
	}
	return sum
}