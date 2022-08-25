package main

import (
	"fmt"
	// "time"
)

// func hello() {
// 	fmt.Println("hello function")
// }
// func main() {
// 	go hello()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("main function")
// }

// func numbers () {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(250 * time.Millisecond)
// 		fmt.Printf("%v",i)
// 	}
// }
// func alphabets () {
// 	for i := 'a'; i <= 'e'; i++ {
// 		time.Sleep(400 * time.Millisecond)
// 		fmt.Printf("%c",i)
// 	}
// }

// func main() {
// 	go numbers()
// 	go alphabets()
// 	time.Sleep(3000 * time.Millisecond)
// 	fmt.Println("\nmain terminated")
// }

// hello function using channel
// func hello(done chan bool){
// 	fmt.Println("hello function")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<- done
// 	fmt.Println("main function")
// }

// hello function using channel and sleep
// func hello(data chan bool) {
// 	fmt.Println("hello is going to sleep")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("sleep is over")
// 	data <- true
// }
// func main() {
// 	data := make(chan bool)
// 	fmt.Println("hello is calling")
// 	go hello(data)
// 	<- data
// 	fmt.Println("main terminated")
// }

// produce function using channel in for range loop
func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
// func main(){
// 	ch := make(chan int)
// 	go produce(ch)
// 	for v := range ch {
// 		fmt.Println("Received: ",v)
// 	}
// }
func main(){
	ch := make(chan int)
	fmt.Println(ch)
	go produce(ch)
	for {
		v, ok := <- ch
		if ok == false {
			fmt.Println("Received : ", v, ok)
			break
		}
		fmt.Println("Received : ", v, ok)
	}
}
