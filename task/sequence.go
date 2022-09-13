package task

import (
	"fmt"
	"sync"
)

type Sequence struct {
	Limit int
}

func (seq *Sequence) OddSequence(wg *sync.WaitGroup) {
	var result []int
	for i := 1; i <= seq.Limit; i += 2 {
		result = append(result, i)
	}
	fmt.Println("Odd Sequence: ", result)
	wg.Done()
}

func (seq *Sequence) EvenSequence(wg *sync.WaitGroup) {
	var result []int
	for i := 2; i <= seq.Limit; i += 2 {
		result = append(result, i)
	}
	fmt.Println("Even Sequence: ", result)
	wg.Done()
}

func (seq *Sequence) PrimeSequence(wg *sync.WaitGroup) {
	var result []int
	for i := 1; i <= seq.Limit; i++ {
		check := 0
		for divider := 1; divider <= i; divider++ {
			if i%divider == 0 {
				check++
			}
		}
		if check == 2 {
			result = append(result, i)
		}
	}
	fmt.Println("Prime Sequence: ", result)
	wg.Done()
}

func (seq *Sequence) FiboSequence(wg *sync.WaitGroup) {
	var result []int
	a, b := 0, 1
	var fibo int
	for a < seq.Limit {
		fibo = a
		a = b
		b = fibo + a
		result = append(result, fibo)
	}
	fmt.Println("Fibo Sequence: ", result)
	wg.Done()
}
