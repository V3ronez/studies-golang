package main

import "fmt"

func main() {
	var array1 [2]int
	array1[0] = 1
	array1[1] = 2

	array2 := [3]int{1, 2, 3}

	array_string := [3]string{"posicao_1", "posicao_2", "posicao_3"}
	fmt.Println(array1, array2, array_string)
	fmt.Println("--------------")
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 10)
	slice_string := []string{"slice1", "slice2"}
	slice2 := array2[0:3] //pega uma fatia do array
	fmt.Println(slice, slice2, slice_string)
}
