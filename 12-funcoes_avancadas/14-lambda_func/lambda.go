package main

import "fmt"

func main() {
	func(text string) {
		fmt.Println("lambda func :). ", text)
	}("esse é o parametro")
}
