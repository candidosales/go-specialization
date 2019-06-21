package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var host chan int
var Chopsticks sync.Pool

type Chopstick struct {
	sync.Mutex
	name int
}

type Philosopher struct {
	leftChopstick, rightChopstick *Chopstick
	name                          int
}

// Func to init pool
func initPool() {
	Chopsticks = sync.Pool{
		New: func() interface{} {
			return new(Chopstick)
		},
	}
}

func (philosopher Philosopher) eat() {
	defer wg.Done()
	//Taking permission from host
	<-host

	philosopher.leftChopstick = Chopsticks.Get().(*Chopstick)
	philosopher.rightChopstick = Chopsticks.Get().(*Chopstick)

	fmt.Printf("starting to eat %d\n", philosopher.name)
	fmt.Printf("finishing eating %d\n", philosopher.name)

	Chopsticks.Put(philosopher.leftChopstick)
	Chopsticks.Put(philosopher.rightChopstick)

	host <- 1
}

func main() {
	host = make(chan int, 2)

	initPool()

	for i := 0; i < 5; i++ {
		chopstick := Chopstick{}
		chopstick.name = i + 1
		Chopsticks.Put(&chopstick)
	}

	Philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		Philosophers[i] = new(Philosopher)
		Philosophers[i].name = i + 1
	}

	for j := 0; j < 3; j++ {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go Philosophers[i].eat()
		}
	}

	host <- 1
	host <- 1

	wg.Wait()
}
