package main

import "fmt"

func Sum(nums ...int) int {
	res := 0

	for _, v := range nums {
		res += v
	}
	return res
}

func ApplyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibnocci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("Hellow World")

	greet := func(msg string) {
		fmt.Println(msg)
	}

	greet("Good Morning...!")

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(ApplyOperation(2, 4, add))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6))

	pos, neg := adder(), adder()

	fmt.Println(pos(1))
	fmt.Println(pos(2))
	fmt.Println(neg(-1))

	fib := fibnocci()

	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
