package main

import (
	"time"
	
	"fmt"
)


func test1(done chan struct{}){
	for{
		select{
		case <-done:
			fmt.Println("test1")
			return
		}
	}
}

func test2(done chan struct{}){
	for{
		select{
		case <- done:
			fmt.Println("test2")
			return
		}
	}
}
func test(){
	done := make(chan struct{})
	defer close(done)
	go test1(done)
	go test2(done)
	
}
func main(){
	test()
	time.Sleep(time.Second)
}