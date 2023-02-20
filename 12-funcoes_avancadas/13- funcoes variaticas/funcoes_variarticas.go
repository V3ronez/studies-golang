package main

import "fmt"

func soma(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
func main() {
	soma(10, 10, 10)
}
