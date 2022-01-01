package main

import "fmt"

func main() {
	Hello()
}

func Hello() {
	fmt.Print("Let's study golang!")
}

func somePrivateMethod() string {
	return "Private Method Call Example"
}
