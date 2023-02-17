package main

import "fmt"

func main() {

	user := map[string]string{
		"nome": "henrique",
	}

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "henrique",
		},
		"dados": {
			"idade": "23",
		},
	}
	fmt.Println("user => ", user, "user 2 =>", user2)
}
