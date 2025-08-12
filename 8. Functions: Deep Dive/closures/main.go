package main

import "fmt"

type TranformFn func(int) int

func main(){
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	doubled := transform(&numbers, genTransformFn(2))
	fmt.Println(doubled)

	tripled := transform(&numbers, genTransformFn(3))
	fmt.Println(tripled)
}

func transform(numbers *[]int, transformFn TranformFn) []int{
	transformedNumbers := []int{}
	for _, val := range *numbers{
		transformedNumbers = append(transformedNumbers, transformFn(val))
	}
	return transformedNumbers
}

func genTransformFn(factor int) TranformFn{
	return func(val int) int{
		return val * factor
	}
}