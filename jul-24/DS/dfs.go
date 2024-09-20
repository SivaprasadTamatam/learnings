package main

import "fmt"

func dfs(g map[int][]int, s int, v map[int]bool) {
	v[s] = true
	fmt.Println(s)

	for _, n := range g[s] {
		if !v[n] {
			dfs(g, n, v)
		}
	}
}

func main() {
	g := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {6},
		6: {},
	}

	v := make(map[int]bool)

	for node := 1; node < 7; node++ {
		fmt.Println("Start - ", node)
		if !v[node] {
			dfs(g, node, v)
		}
	}

}
