package main

import (
	"fmt"
	"sync"
)

func printMsg(text string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(text)
}

func main(){
	var wg sync.WaitGroup

	wg.Add(3)
	go printMsg("This message is from goroutine 1", &wg)
	go printMsg("This message is from goroutine 2", &wg)
	go printMsg("This message is from goroutine 3", &wg)

	fmt.Println("Loading...")
	wg.Wait()
	fmt.Println("All goroutine has be done!")
}