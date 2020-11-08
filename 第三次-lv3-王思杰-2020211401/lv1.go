package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var lock sync.Mutex
func f1(a chan int){
	defer wg.Done()
	for i:=0;i<=100;i++{
		a <- i
		fmt.Println(i)
		}
	}
func f2(a chan int){
	defer wg.Done()
	for i:=0;i<=100;i++{
		<-a
		fmt.Println(i)}
}


func main(){
	msg :=make(chan int)
	wg.Add(2)
	go f1(msg)
	go f2(msg)

wg.Wait()
}