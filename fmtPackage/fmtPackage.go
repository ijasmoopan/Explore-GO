package main

import "fmt"

func main() {

	p := fmt.Println

	a := fmt.Sprint("Hello", "Hi")
	b := fmt.Sprint("World")
	p(a)
	p(b)
	c := fmt.Sprintln("Hello","World")
	p(c)
	fmt.Print(c)
	v := 10
	d := fmt.Sprintf("Hi %v",v)
	p(d)
	var name string
	p("Your name please : ")
	fmt.Scan(&name)
	fmt.Println("Nice to meet you", name)
}