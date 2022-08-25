package main

import (
	"fmt"
	// "time"
)

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "ijas"
// 	ch <- "moopan"
// 	// ch <- "mo"
// 	fmt.Println(<- ch)
// 	fmt.Println(<- ch)
// 	// fmt.Println(<- ch)
// }

// func write (ch chan int) {
// 	for i := 0; i < 5; i ++ {
// 		ch <- i
// 		fmt.Println("Succesfully wrote : ", i, "===================")
// 	}
// 	close(ch)
// }
// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	fmt.Println("main is going to sleep")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("sleep is over")
// 	for v := range ch {
// 		fmt.Println("Received value : ", v, "++++++++++++++++++++")
// 		fmt.Println("main is going to another sleep")
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("another sleep is over")
// 	}
// }

// func main() {
// 	ch := make(chan int, 5)
// 	ch <- 5
// 	ch <- 6
// 	ch <- 7
// 	ch <- 8
// 	ch <- 9
// 	// ch <- 10
// 	close(ch)

// 	n, open := <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// 	n, open = <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// 	n, open = <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// 	n, open = <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// 	n, open = <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// 	n, open = <- ch
// 	fmt.Printf("recieved : %d open : %t \n",n, open )
// }

// func main() {
// 	ch := make (chan int, 5)
// 	ch <- 5
// 	ch <- 6
// 	close(ch)
// 	fmt.Println("Capacity : ",cap(ch))
// 	fmt.Println("Length : ", len(ch))
// 	// for v := range ch {
// 	// 	fmt.Printf("recieved : %d \n",v)
// 	// }
// 	fmt.Println("Received : ", <- ch)
// 	fmt.Println("Capacity : ",cap(ch))
// 	fmt.Println("Length : ", len(ch))
// }

func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Wrote: ", i)
		// fmt.Printf("ch: %v\n", ch)
	}
}
func main() {
	ch := make(chan int, 2)
	go sendData(ch)
	for i := 0; i < 5; i++ {
		a := <-ch
		fmt.Println("Read: ", a)
	}
}
