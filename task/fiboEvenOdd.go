package task

import (
	"fmt"
	"sync"
)

func Fibonacci(num int, ch chan<- []int, wg *sync.WaitGroup) {
	var result []int
	a, b := 0, 1
	var fibo int
	for a < num {
		fibo = a
		a = b
		b = fibo + a
		result = append(result, fibo)
	}
	ch <- result
	wg.Done()
}

func EvenOdd(ch <-chan []int, wg *sync.WaitGroup) {
	for _, v := range <-ch {
		if v%2 == 0 {
			fmt.Println(v, " is Even number")
		} else {
			fmt.Println(v, " is Odd number")
		}
	}
	wg.Done()
}
