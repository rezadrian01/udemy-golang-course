package main

import "fmt"

type TransformFn func(int) int

func main(){
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func (val int) int{ 
		return val * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform TransformFn) []int{
	dNumbers := []int{}
	for _, val := range *numbers{
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}