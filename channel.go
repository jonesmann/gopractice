package main

import (
	"fmt"
)

type NonBlock struct{
	cin chan string
	con chan string
}

func NewNonBlock()*NonBlock{
	nbc := NonBlock{make(chan string), make(chan string)}
	var v string
	go func(){
		for{
			select{
				case v = <- nbc.cin:
				case nbc.con <- v: 
			}
		}
	}()
	return &nbc
}
func (nbc *NonBlock)Send(v string){
	nbc.cin <- v
}
func (nbc *NonBlock)Recieve()string{
	return <-nbc.con
}

func main(){
	nbc := NewNonBlock()
	stopflag := make(chan bool)
	go func(){
		for i:=0;i<10;i++{
			nbc.Send(fmt.Sprint(i))
		}
		stopflag <- true
	}()
	<- stopflag
	fmt.Println("Recieve value ", nbc.Recieve())	
}