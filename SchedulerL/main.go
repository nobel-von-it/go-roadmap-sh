package main

import (
	"log"
	"runtime"
)

func main() {
	example()
}

func example() {
	var x int
	delete(map[int]int{}, x)
	log.Println(runtime.NumCPU())
}

// go tool objdump -s -l example.o
