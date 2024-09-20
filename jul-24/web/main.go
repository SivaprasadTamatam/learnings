package main

import "sort"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func MakeChange(amount int) map[byte]int {
	coins := map[byte]int{'H': 50, 'Q': 25, 'D': 10, 'N': 5, 'P': 1}
	change := map[byte]int{}

	var coinValues []int

	mapValuetoSym := make(map[int]byte)

	for coin, val := range coins {
		mapValuetoSym[val] = coin
		coinValues = append(coinValues, val)
	}

	sort.Slice(coinValues, func(i, j int) bool {
		return coinValues[i] > coinValues[j]
	})

	for _, v := range coinValues{
		if amount >= v {
			quot, _ := divmod(amount, v)
			change[mapValuetoSym[v]] = quot
			amount -= quot *v
			
		}
	}

	/*for coin, val := range coins {

		if amount >= val {
			quot, _ := divmod(amount, val)
			change[coin] = quot
			amount -= quot * val
			change[coin] = quot
		}

	}*/

	return change
}
