package main

import (
	"fmt"
	"time"
)


func goro_print(number string){
	for i:=0; i<1; i++ {
		fmt.Printf("I am %s goro\n", number)
		time.Sleep(time.Millisecond)
	}
}


// go 协程体验

func main(){
	go goro_print("1")
	go goro_print("2")
	go goro_print("3")

	time.Sleep(time.Second)
}