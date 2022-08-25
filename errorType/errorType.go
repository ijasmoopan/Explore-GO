package main

import "fmt"

type MyError struct {
	errorType string
}

func (myErr *MyError) Error() string {
	if myErr.errorType == "network timeout" {
		return "ERROR: No internet access, please check your Wi-Fi connection"
	} else if myErr.errorType == "invalid creds" {
		return "ERROR: Incorrect username or password provided"
	} else {
		return "ERROR: Encountered an unknown error"
	}

}

func main() {
	var myErr1 *MyError = nil

	myErr2 := MyError{
		errorType: "invalid creds",
	}
	myErr3 := &MyError{
		errorType: "invalid creds",
	}
	fmt.Println(myErr1)
	fmt.Println(myErr2)
	fmt.Println(myErr3)

	myErr4 := &MyError {
		errorType: "",
	}
	fmt.Println(myErr4)
}