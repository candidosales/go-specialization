/* Implement the dining philosopher’s problem with the following constraints/modifications.

> There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
> Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
> The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
> In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
> The host allows no more than 2 philosophers to eat concurrently.
> Each philosopher is numbered, 1 through 5.
> When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
> When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type chStick struct {
}
type philfn struct {
	phNum                   int
}

var heapCS = sync.Pool{
	New:func() interface{} {
		return new(chStick)
	},
}

func (p philfn) eat(count chan int) {
	defer wg.Done()
	<-count
	left := heapCS.Get()
	right := heapCS.Get()
	fmt.Println("Philosopher", p.phNum, "is eating")//, x, "times")
	time.Sleep(time.Second)
	heapCS.Put(left)
	heapCS.Put(right)
	fmt.Println("Philosopher", p.phNum, "finished eating")//, x, "times")
	time.Sleep(time.Second)
	count<- 1
}

func policeRoutine(count chan int) {
	defer wg.Done()
	count <- 1
	//count <- 1
}

var wg sync.WaitGroup

func main() {
	count := make(chan int, 2)
	for i := 0; i < 5; i++ {
		heapCS.Put(new(chStick))
	}
	philo := make([]*philfn, 5)
	for i := 0; i < 5; i++ {
		philo[i] = &philfn{i+1}
		for j:=0;j<3;j++ {
			wg.Add(1)
			//fmt.Println(i)
			go philo[i].eat(count)
		}
	}
	wg.Add(1)
	go policeRoutine(count)
	wg.Wait()
}
