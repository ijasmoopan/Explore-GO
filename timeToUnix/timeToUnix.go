package main

import (
	"fmt"
	// "strconv"

	// "os"
	"time"
)

// func main() {
// 	var t time.Time
// 	t = time.Now() //It will return time.Time object with current timestamp
// 	fmt.Println("Time : ", t)

// 	tUnix := t.Unix() //Converstion to Time Unix also known as epoch time
// 	fmt.Println("Time Unix : ", tUnix)

// 	//Conversion to Time Unix Millisecond
// 	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
// 	fmt.Println("Time Unix Milli : ", tUnixMilli)

// 	//Conversion to Time Unix Microsecond
// 	tUnixMicro := int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
// 	fmt.Println("Time Unix Micro : ", tUnixMicro)

// 	//Conversion to Time Unix Nanosecond
// 	tUnixNano := t.UnixNano()
// 	fmt.Println("Time Unix Nano : ", tUnixNano)
// }

// func main() {
// 	tUnix := time.Now().Unix()
// 	fmt.Println("Time Unix : ", tUnix)
// }

// func main() {
// 	dateString := "2021-11-22"
// 	date, error := time.Parse("2006-01-02", dateString)
// 	if error != nil {
// 		fmt.Println(error)
// 		return
// 	}
// 	fmt.Printf("Type of dateString: %T\n", dateString)
// 	fmt.Printf("Type of date: %T\n", date)
// 	fmt.Println()
// 	fmt.Printf("Value of dateString: %v\n", dateString)
// 	fmt.Printf("Value of date: %v", date)
// }

// func main() {
// 	timeStamp := int64(1585990763)
// 	date := time.Unix(timeStamp, 0)
// 	fmt.Println(date)

// 	thetime, e := time.Parse(time.RFC3339, "2020-04-04T22:08:41+00:00")
// 	if e != nil {
// 		panic("Can't parse time format")
// 	}
// 	fmt.Println(thetime)
// 	epoch := thetime.Unix()
// 	fmt.Fprintf(os.Stdout, "Epoch: %d\n", epoch)
// }


// const(
// 	layoutISO = "2006-01-02"
// 	layoutUS = "January 2, 2006"
// 	Clock = "14:24:30"
// 	timeLayout = "3:04:05 PM"
// )
// func main() {
// 	date := "2020-04-06"
// 	t , err := time.Parse(layoutISO, date)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(t)
// 	t2 := t.Format(layoutUS)
// 	fmt.Println(t2)

// 	hour, minute, second := time.Now().Clock()
// 	tym := fmt.Sprint(strconv.Itoa(hour),":", strconv.Itoa(minute),":", strconv.Itoa(second))
// 	fmt.Println("tym : ",tym)
// 	tym1, err := time.Parse(Clock, tym)
// 	if err != nil {
// 		fmt.Println("error : ",err)
// 	}
// 	fmt.Println("tym1 : ",tym1)
// 	tym2 := tym1.Format(timeLayout)
// 	fmt.Println("tym2 : ",tym2)

// 	a := "12 45 30"
// 	b, err := time.Parse(Clock, a)
// 	if err != nil {
// 		fmt.Println("error : ", err)
// 	}
// 	fmt.Println("b : ",b)
// 	c := b.Format(timeLayout)
// 	fmt.Println("c : ",c)
// }


func main() {
  
    // this function returns the present time
    current_time := time.Now()
  
    // individual elements of time can
    // also be called to print accordingly
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
    current_time.Year(), current_time.Month(), current_time.Day(),
    current_time.Hour(), current_time.Minute(), current_time.Second())

    // formatting time using
    // custom formats
    fmt.Println(current_time.Format("2006-01-02 15:04:05"))
    fmt.Println(current_time.Format("2006-January-02"))
    fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))
}
