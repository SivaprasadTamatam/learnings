package main

import "fmt"

func main() {

	var i interface{} = "Hello"

	switch v := i.(type) {
	case string:
		fmt.Println(v)
	}

}
