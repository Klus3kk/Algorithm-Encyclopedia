package main

import "fmt"


func nwd(a, b int) int {
	for b != 0 && a != 0 { 
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}

var a = 5
var b = 1

func main() {
	fmt.Println(nwd(a,b))
}