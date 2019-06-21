package main

import (
	"fmt"
	"sync"
	"time"
)

// Implement the dining philosopher’s problem with the following constraints/modifications:
// 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// 5. The host allows no more than 2 philosophers to eat concurrently.
// 6. Each philosopher is numbered, 1 through 5.
// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number int
	count  int

	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (philosopher Philosopher) eat(channel chan *Philosopher, waitGroup *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		channel <- &philosopher

		if philosopher.count < 3 {
			philosopher.leftChopstick.Lock()
			philosopher.rightChopstick.Lock()

			fmt.Println("starting to eat ", philosopher.number)
			philosopher.count++
			fmt.Println("finishing eating ", philosopher.number)

			philosopher.leftChopstick.Unlock()
			philosopher.rightChopstick.Unlock()

			waitGroup.Done()
		}
	}
}

func host(channel chan *Philosopher, waitGroup *sync.WaitGroup) {
	for {
		if len(channel) == 2 {
			<-channel
			<-channel

			// delay
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var waitGroup sync.WaitGroup
	philosophersChannel := make(chan *Philosopher, 2)

	waitGroup.Add(15)

	chopsticks := make([]*Chopstick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, 0, chopsticks[i], chopsticks[(i+1)%5]}
	}

	go host(philosophersChannel, &waitGroup)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(philosophersChannel, &waitGroup)
	}

	waitGroup.Wait()
}
