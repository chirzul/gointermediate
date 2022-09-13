package main

import (
	"fmt"
	"gointermediate/task"
	"sync"
)

var wg = &sync.WaitGroup{}
var mt = &sync.RWMutex{}

func main() {
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)
	defer close(chn)
	wg.Add(2)
	go task.Sum(a[:len(a)/2], chn, wg)
	go task.Sum(a[len(a)/2:], chn, wg)
	//receive
	res1 := <-chn
	fmt.Println("hasil:", res1)
	res2 := <-chn
	fmt.Println("hasil:", res2)
	wg.Wait()

	sequence := task.Sequence{Limit: 40}
	wg.Add(4)
	go sequence.OddSequence(wg)
	go sequence.EvenSequence(wg)
	go sequence.PrimeSequence(wg)
	go sequence.FiboSequence(wg)
	wg.Wait()

	chEvenOdd := make(chan []int)
	defer close(chEvenOdd)
	wg.Add(2)
	go task.Fibonacci(40, chEvenOdd, wg)
	go task.EvenOdd(chEvenOdd, wg)
	wg.Wait()

	chPoke := make(chan string, 5)
	defer close(chPoke)
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		mt.Lock()
		go venusaur.AddBet(wg, mt)
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		mt.Lock()
		go charizard.AddBet(wg, mt)
	}
	chBet := make(chan float64)
	defer close(chBet)
	mt.RLock()
	go venusaur.GetBet(chBet, wg, mt)
	fmt.Println("Bet for", venusaur.Name, ":", venusaur.Bet)
	mt.RLock()
	go charizard.GetBet(chBet, wg, mt)
	fmt.Println("Bet for", charizard.Name, ":", charizard.Bet)
	wg.Add(1)
	go task.Tournament(venusaur, charizard, 3, chPoke, wg)
	fmt.Println(<-chPoke)
	wg.Wait()
}

var venusaur = &task.Pokemon{
	Name: "Venusaur",
	Hp:   100,
	Skills: []task.Skill{
		{Name: "Basic Attack", Damage: 5},
		{Name: "Vine Whip", Damage: 15},
		{Name: "Solar Beam", Damage: 25},
	},
}

var charizard = &task.Pokemon{
	Name: "Charizard",
	Hp:   100,
	Skills: []task.Skill{
		{Name: "Basic Attack", Damage: 5},
		{Name: "Wing Attack", Damage: 15},
		{Name: "Fire Blast", Damage: 25},
	},
}
