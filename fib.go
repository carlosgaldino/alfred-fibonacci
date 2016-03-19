package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(n int) int {
	a, b := 0, 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := fib(r1.Intn(7))
	item := fmt.Sprintf("<item arg ='%d' uid ='%d'>\n<title>%d</title>\n<subtitle>%d</subtitle>\n</item>", n, n, n, n)
	fmt.Printf("<?xml version='1.0'?>\n<items>\n%s</items>", item)
}
