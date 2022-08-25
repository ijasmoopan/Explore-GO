package main

import (
	"fmt"
	"time"
)

// func server1(output1 chan string) {
// 	time.Sleep(2 * time.Second)
// 	output1 <- "from server1"
// }
// func server2(output2 chan string) {
// 	time.Sleep(3 * time.Second)
// 	output2 <- "from server2"
// }
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)
// 	go server1(output1)
// 	go server2(output2)
// 	select {
// 	case s1 := <- output1 :
// 		fmt.Println(s1)
// 	case s2 := <- output2 :
// 		fmt.Println(s2)
// 	}
// }

// --------------------

func process(ch chan string) {  
    time.Sleep(2500 * time.Millisecond)
    ch <- "process successful"
}

func main() {  
    ch := make(chan string)
    go process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default:
            fmt.Println("no value received")
        }
    }

}