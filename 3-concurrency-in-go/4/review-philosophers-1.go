package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }
type Philosopher struct {
	Id               int
	Left, Right      *Chopstick
	permission, done chan<- interface{}
}

func (p *Philosopher) eat() {
	for i := 0; i < 3; i++ {

		p.permission <- p.Id

		leftFirst := rand.Float32() > 0.5
		if leftFirst {
			p.Left.Lock()
			p.Right.Lock()
		} else {
			p.Right.Lock()
			p.Left.Lock()
		}

		fmt.Printf("%sstarting to eat %d\n", strings.Repeat(" ", 5-p.Id), p.Id)
		time.Sleep(time.Millisecond * time.Duration(500-rand.Intn(500)))
		fmt.Printf("%sfinished eating %d\n", strings.Repeat(" ", 5-p.Id), p.Id)

		if leftFirst {
			p.Left.Unlock()
			p.Right.Unlock()
		} else {
			p.Right.Unlock()
			p.Left.Unlock()
		}

		p.done <- p.Id
	}
}

func host(max int, permission <-chan interface{}, done <-chan interface{}) {
	count := 0
	for {
		select {
		case <-done:
			count--
			//fmt.Printf("%d is done\n", id)
		default:
			if count < max {
				<-permission
				count++
				//fmt.Println("granted permission to", id)
			}
		}
	}
}

var wg sync.WaitGroup

func main() {

	permission := make(chan interface{})
	done := make(chan interface{}, 1)

	// initialize the sticks and philosophers
	sticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			Id:         i + 1,           // Id
			Left:       sticks[i],       // Left
			Right:      sticks[(i+1)%5], // Right
			permission: permission,
			done:       done,
		}
	}

	// let all philosophers eat

	max := 2

	go host(max, permission, done)

	for _, p := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			p.eat()
			wg.Done()
		}(p)
	}

	wg.Wait()
	close(done)
}
