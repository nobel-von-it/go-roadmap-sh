package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

func Eq[T comparable](a, b T) bool {
	return a == b
}

func Cmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func main() {
	var a, b int
	if c, err := fmt.Scanf("%d %d", &a, &b); err != nil && c != 2 {
		log.Fatal(err)
	}
	fmt.Println(Eq(a, b))
	fmt.Println(Cmp(a, b))
}
