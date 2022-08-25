package main

import "log"

func main() {
	log.Fatal("Exception occured!")
	s := "formatted"
	log.Fatalf("A %s string for logging", s)
	log.Fatalln("A logged string with a new line!")
}