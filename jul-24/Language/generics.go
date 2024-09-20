package main

import "fmt"

func Add(nums ...any) (result float64) {

	for _, num := range nums {
		switch v := num.(type) {
		case int:
			result += float64(v)
		case float64:
			result += v
		default:
			return 0

		}
	}
	return result
}

func main() {
	fmt.Println(Add(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(Add(1.2, 3.4, 5.4, 8.9))
}
