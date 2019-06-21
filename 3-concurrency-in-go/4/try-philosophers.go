package main

import (
	"fmt"
	"sync"
)

const count = 5
const meals = 3

func main() {
	sticks := make([]*ChopS, count)
	for i := 0; i < count; i++ {
		sticks[i] = new(ChopS)
	}

	philos := make([]*Philo, count)
	for i := 0; i < count; i++ {
		left := i
		right := (i + 1) % count
		min, max := minMax(left, right)
		fmt.Printf("Philo %d takes >> left: %d, right: %d\n", i, min, max)
		philos[i] = &Philo{i + 1, meals, sticks[min], sticks[max]}
	}

	ch := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go host(ch, wg)

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

func host(ch chan int, wg sync.WaitGroup) {
	println("Started host.")
	for i := 0; i < count*meals; i++ {
		<-ch
	}
	wg.Done()
}

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id, meals       int
	leftCS, rightCS *ChopS
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
