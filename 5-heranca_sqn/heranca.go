package main

import "fmt"

type people struct {
	name string
	age  uint8
}

type student struct {
	id uint8
	people
}

func main() {

	user := people{name: "henrique", age: 10}
	student := student{23, user}
	fmt.Println(student.name)
}
