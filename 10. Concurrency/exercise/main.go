package main

import (
	"fmt"
	"time"
)


func greet(text string, doneChan chan bool){
	fmt.Println(text)
	doneChan <- true
}

func slowGreet(text string, doneChan chan bool){
	time.Sleep(3 * time.Second)
	fmt.Println(text)

	doneChan <- true
	close(doneChan)
}

func main(){
	doneChans := make([]chan bool, 4)
	for index := range doneChans{
		doneChans[index] = make(chan bool)
	}
	go greet("Hello Ahmad Reza Adrian", doneChans[0])
	go greet("How are you", doneChans[1])
	go slowGreet("This take long time", doneChans[2])
	go greet("End", doneChans[3])
	
	for index := range doneChans{
		<- doneChans[index]
	}
}