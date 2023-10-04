package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
var a = 0
var b = 1


func fibonacci() func() int {
	return func () (c int) {
		c = a + b
		a = b
		b = c
		return 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
