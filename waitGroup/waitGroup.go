package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Goroutine ", i)
	time.Sleep(3 * time.Second)
	fmt.Println("Goroutine ended : ", i)
	wg.Done()
}
func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go produce(i, &wg)
	}
	fmt.Println("Calling wait function..")
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}
