package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())  // It prints the actual time
	// time.Sleep(2 * time.Second)  // It sleeps the go thread for 2 second
	// fmt.Println(time.Now()) // It prints the actual time after that 2 second

	// Duration
	// t1 := time.Now()
	// time.Sleep(2 * time.Second)
	// t2 := time.Now()
	// fmt.Println("Duration was : ",t2.Sub(t1)) // It prints the duration of t1 and t2

	// t1 := time.Now()
	// time.Sleep(3 * time.Second)
	// t2 := time.Now()
	// t := t2.Sub(t1)
	// fmt.Println(t, t.Milliseconds(), t.Microseconds(), t.Nanoseconds())


	// Full time details..
	var t time.Time
	t = time.Now()

	fmt.Println(t)
	fmt.Println(t.Date())
	fmt.Println(t.Day())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Zone()) // Time zone
	fmt.Println(t.UnixNano())
	fmt.Println(t.Clock())


	// The before and after functions check whether a time is before or after another time.
	t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.After(t2)) // false
}