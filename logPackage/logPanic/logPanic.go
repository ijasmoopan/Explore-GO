package main

import "log"

func main() {
	// log.Panic("Error!")
	// panic("Error")

	log.Println("Print")

	// s, v := "formatted", 42
	// log.Panicf("A %s error with num %d", s, v)

	// log.Panicln("A panic with a new line!")

	log.SetFlags(0)
	log.Println("Print")

	log.SetFlags(1)
	log.Println("Print")

	log.SetFlags(2)
	log.Println("Print")

	log.SetFlags(3)
	log.Println("Print")

	log.SetFlags(2 | 3)
	log.Println("Print")
}