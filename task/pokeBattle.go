package task

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pokemon struct {
	Name   string
	Hp     int
	Skills []Skill
	Bet    float64
}
type Skill struct {
	Name   string
	Damage int
}

func Battle(p1 *Pokemon, p2 *Pokemon, ch chan<- string) {
	var p1Skill = p1.Skills[rand.Intn(len(p1.Skills))]
	var p2Skill = p2.Skills[rand.Intn(len(p2.Skills))]
	p1.Hp -= p2Skill.Damage
	p2.Hp -= p1Skill.Damage
	fmt.Println("      ___---Bak buk Bak buk---___")
	fmt.Printf("%s launch %s! deals %d!\n", p1.Name, p1Skill.Name, p1Skill.Damage)
	fmt.Printf("%s launch %s! deals %d!\n", p2.Name, p2Skill.Name, p2Skill.Damage)
	fmt.Printf("%s hp: %d, %s hp: %d\n", p1.Name, p1.Hp, p2.Name, p2.Hp)

	time.Sleep(1 * time.Second)
	if p1.Hp > 0 && p2.Hp > 0 {
		Battle(p1, p2, ch)
	} else if p1.Hp > 0 {
		fmt.Println(p1.Name, "menang")
		ch <- "p1"
	} else if p2.Hp > 0 {
		fmt.Println(p2.Name, "menang")
		ch <- "p2"
	} else {
		fmt.Println("Pertarungan seri.")
		ch <- "seri"
	}
}

func Tournament(p1 *Pokemon, p2 *Pokemon, match int, ch chan string, wg *sync.WaitGroup) {
	rand.Seed(time.Now().Unix())
	// var tournamentResult string
	chBattle := make(chan string)
	defer close(chBattle)
	var p1Score, p2Score int
	for i := 1; i <= match; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Memulihkan hp %s dan %s\nBattle ke-%d start!\n", p1.Name, p2.Name, i)
		p1.Hp = 100
		p2.Hp = 100
		time.Sleep(1 * time.Second)
		go Battle(p1, p2, chBattle)
		battleResult := <-chBattle
		if battleResult == "p1" {
			p1Score++
		} else if battleResult == "p2" {
			p2Score++
		}
		fmt.Println("Skor", p1Score, p2Score)
	}
	if p1Score > p2Score {
		reward := (p1.Bet + p2.Bet) / p1.Bet * 5000
		ch <- fmt.Sprintf("%s Memenangkan turnamen\nPara supporter mendapatkan return %f", p1.Name, reward)
	} else if p2Score > p1Score {
		reward := (p1.Bet + p2.Bet) / p2.Bet * 5000
		ch <- fmt.Sprintf("%s Memenangkan turnamen\nPara supporter mendapatkan return %f", p2.Name, reward)
	} else {
		ch <- "Turnamen berakhir seri!"
	}
	wg.Done()
}

func (p *Pokemon) AddBet(wg *sync.WaitGroup, mt *sync.RWMutex) {
	p.Bet++
	mt.Unlock()
	wg.Done()
}

func (p *Pokemon) GetBet(ch chan float64, wg *sync.WaitGroup, mt *sync.RWMutex) {
	bet := p.Bet
	mt.RUnlock()
	ch <- bet
	wg.Done()
}
