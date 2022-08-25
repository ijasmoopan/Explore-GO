package main

import (
	"fmt"
	// "strings"
)

func main() {
	var s string
	s = "hello, world"
	fmt.Println(s)
	fmt.Println(len(s)) 

	// Uncomment the following line to change a value
  	// in the string. An error will be thrown.
  	// s[1] = "c"

	fmt.Println(s)
	for index, char := range s {
		fmt.Println("index : ", index, "character : ", char)
		fmt.Println("index : ", index, "character : ", string(char))
		fmt.Printf("index : %d  character : %c \n", index, char)
	}

	slice := []byte {0x48, 0x65, 0x6C,  0x6C, 0x6f} 
	str := string(slice)
	fmt.Println(str)

	data := "Netflix"
	fmt.Println(mutate([]rune(data)))
	
	stringLiterals()
}
func mutate (str []rune) string {
	str[0] = 'M'
	return string(str)
}

func stringLiterals () {
	str := `Hello
	world`
	fmt.Println(str)
}