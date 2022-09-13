package task

import (
	"fmt"
	"sync"
)

func Sum(d []int, ch chan int, wg *sync.WaitGroup) {
	var result int
	for _, v := range d {
		//hitung
		fmt.Print(v, " ")
		result += v
	}
	// send result
	ch <- result
	wg.Done()
}
