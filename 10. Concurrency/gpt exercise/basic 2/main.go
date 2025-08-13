package main

import "fmt"

func main(){
	ch := make(chan int, 3)
	fmt.Println("Chan capacity: ", cap(ch))

	fmt.Println("Sending 1, 2, 3 to chanels...")
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	fmt.Println("Numbers has been sended!")
	for num := range ch{
		fmt.Println(num)
	}
}