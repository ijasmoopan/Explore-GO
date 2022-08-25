package main

import (
	"errors"
	"fmt"
)

func main() {

	// err := foo()
	// if err != nil {
	// 	fmt.Println("Something error occured")
	// 	return err
	// } else {
	// 	fmt.Println("No error")
	// }

	result, err := lessThanTen(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result, err = lessThanTen(20)
	if err != nil {
		fmt.Println(result, err)
		return
	}
	fmt.Println(result)
}

func lessThanTen (i int) (string, error) {
	if i < 10 {
		return "Woohoo..This integer is less than 10", nil
	}
	return "This number is too big", errors.New("ERROR : Invalid value provided")
}