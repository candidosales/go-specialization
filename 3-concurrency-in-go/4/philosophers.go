package main

import (
	"fmt"
	"sync"
)

const chopSticks = 5
const maxMeals = 3

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	id, meals       int
	leftCS, rightCS *ChopStick
}

func (p Philo) eat(host chan int) {
	for i := 0; i < p.meals; i++ {
		host <- p.id

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {

	sticks := make([]*ChopStick, chopSticks)
	for i := 0; i < chopSticks; i++ {
		sticks[i] = new(ChopStick)
	}

	philos := make([]*Philo, chopSticks)
	for i := 0; i < chopSticks; i++ {
		left := i
		right := (i + 1) % chopSticks
		min, max := minMax(left, right)
		fmt.Printf("Philo %d takes >> left: %d, right: %d\n", i, min, max)
		philos[i] = &Philo{i + 1, maxMeals, sticks[min], sticks[max]}
	}

	ch := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go permisssionHost(ch, wg)

	for _, p := range philos {
		go p.eat(ch)
	}

	wg.Wait()
}

func minMax(left int, right int) (int, int) {
	if left < right {
		return left, right
	} else {
		return right, left
	}
}

func permisssionHost(ch chan int, wg sync.WaitGroup) {
	println("Started host.")
	for i := 0; i < chopSticks*maxMeals; i++ {
		<-ch
	}
	wg.Done()
}
