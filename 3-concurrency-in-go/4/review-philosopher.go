package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func philosopher(id int, left *sync.Mutex, right *sync.Mutex, in chan int, out chan int) {
	for i := 0; i < 3; i++ {
		<-in
		left.Lock()
		right.Lock()
		fmt.Println("starting to eat " + strconv.Itoa(id))
		fmt.Println("finishing eating " + strconv.Itoa(id))
		right.Unlock()
		left.Unlock()
		out <- 0
	}
	wg.Done()
}

func host(out []chan int, in []chan int, abort chan int) {
	out[0] <- 0
	for {
		select {
		case out[0] <- 0:
		case out[1] <- 0:
		case out[2] <- 0:
		case out[3] <- 0:
		case out[4] <- 0:
		case <-abort:
		}

		select {
		case <-in[0]:
		case <-in[1]:
		case <-in[2]:
		case <-in[3]:
		case <-in[4]:
		}
	}
}

func main() {
	locks := make([]sync.Mutex, 5)
	host2phi := make([]chan int, 5)
	phi2host := make([]chan int, 5)

	for i := 0; i < 5; i++ {
		host2phi[i] = make(chan int)
		phi2host[i] = make(chan int, 1)
	}

	abort := make(chan int)
	wg.Add(5)
	go host(host2phi, phi2host, abort)
	for i := 0; i < 5; i++ {
		go philosopher(i+1, &locks[i%5], &locks[(i+1)%5], host2phi[i], phi2host[i])
	}

	wg.Wait()
	abort <- 0
}
