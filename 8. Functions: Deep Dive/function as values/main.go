package main

import "fmt"


type TransformFn func(int) int

func main(){
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	doubleNumbers := transform(&numbers, double)
	tripleNumbers := transform(&numbers, triple)

	fmt.Println(doubleNumbers)
	fmt.Println(tripleNumbers)
}

func transform(numbers *[]int, transformFn TransformFn) []int {
	transformedNumbers := []int{}
	for _, val := range *numbers{
		transformedNumbers = append(transformedNumbers, transformFn(val))
	}
	return transformedNumbers
}

func double(val int) int{
	return val * 2
}

func triple(val int) int{
	return val * 3
}