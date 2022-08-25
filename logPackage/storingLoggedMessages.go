package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("ERROR: Can't divide by zero")
	}
	return x / y, nil
}

func main() {	
	var x float64 = 10
	var y float64 = 5

	// Printing log
	// result, err := division(x, y)
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println(result)

	// Storing the logged messages in a file
	result, exception := division(x, y)
	
	file, err := os.OpenFile("info.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print(result)
	log.Print(exception)
	fmt.Println(result)
}