package main

import (
	"errors"
	"fmt"
	"sync"
)

// this is no more required, just keeping it as it is
func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func findAllOddPalindromes(s string) []string {
	n := len(s)
	res := []string{}
	mp := make(map[string]bool)
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			str := s[l : r+1]
			if len(str)%2 == 1 && !mp[str] {
				mp[str] = true
				res = append(res, str)
			}
			l--
			r++
		}

	}
	return res
}

type Result struct {
	val interface{}
	err error
}

func execute(maxThreads int, funcs []func() (interface{}, error)) (interface{}, error) {
	// buffer channel - maxThreads
	// loop the funcs and try to call concurretly
	// waitgroups
	// res

	buffCh := make(chan int, maxThreads)
	res := make(chan Result, len(funcs))
	wg := sync.WaitGroup{}
	done := make(chan bool)

	for _, fn := range funcs {

		select {
		case <-done:
			break
		default:
			buffCh <- 1
			wg.Add(1)
			go func(fn func() (interface{}, error)) {
				defer func() {
					<-buffCh
					wg.Done()
				}()

				val, err := fn()
				if err == nil {
					select {
					case done <- true:
					default:
					}
				}
				res <- Result{val, err}
			}(fn)
		}
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		if r.err == nil {
			return r.val, nil
		}
	}
	return nil, errors.New("no successful execution")
}

func Err_Func() (interface{}, error) {
	return -1, errors.New("execution failed")
}

func Succ_Func() (interface{}, error) {
	return 1, nil
}

func main() {

	fmt.Println(findAllOddPalindromes("a"))     //=> a
	fmt.Println(findAllOddPalindromes("ab"))    //=> a.b
	fmt.Println(findAllOddPalindromes("aba"))   //=> a, b , aba
	fmt.Println(findAllOddPalindromes("xabay")) //=> x, a, b, y, aba
	fmt.Println(findAllOddPalindromes("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

	// all failed ones
	err_funcs := []func() (interface{}, error){Err_Func, Err_Func, Err_Func, Err_Func, Err_Func, Err_Func}
	fmt.Println(execute(3, err_funcs)) // max threads less than no of funcs

	// all failed ones
	err_funcs = []func() (interface{}, error){Err_Func, Err_Func, Err_Func, Err_Func}
	fmt.Println(execute(8, err_funcs)) // max threads greater than no of funcs

	// all success ones
	succ_funcs := []func() (interface{}, error){Succ_Func, Succ_Func, Succ_Func, Succ_Func, Succ_Func, Succ_Func}
	fmt.Println(execute(3, succ_funcs)) // max threads less than no of funcs

	// all success ones
	succ_funcs = []func() (interface{}, error){Succ_Func, Succ_Func, Succ_Func, Succ_Func, Succ_Func, Succ_Func}
	fmt.Println(execute(8, succ_funcs)) // max threads greater than no of funcs

	// mixed  ones
	mixed_funcs := []func() (interface{}, error){Err_Func, Err_Func, Err_Func, Succ_Func, Succ_Func, Succ_Func}
	fmt.Println(execute(3, mixed_funcs)) // max threads less than no of funcs

	// mixed  ones
	mixed_funcs = []func() (interface{}, error){Err_Func, Err_Func, Err_Func, Succ_Func, Succ_Func, Succ_Func}
	fmt.Println(execute(8, mixed_funcs)) // max threads greater than no of funcs
}

//tim.roche@weather.com
