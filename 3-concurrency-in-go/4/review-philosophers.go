package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Dining philosopher’s problem with the following constraints/modifications
-------------------------------------------------------------------------

	1.- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

	2.- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

	3.- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

	4.- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

	5.- The host allows no more than 2 philosophers to eat concurrently.

	6.- Each philosopher is numbered, 1 through 5.

	7.- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

	8.- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

// Constant to define number of philosophers and chopsticks (5)
const numEntities int = 5

// Constant to define the amount of meals for each philosopher (3)
const meals int = 3

// Type chopstick (Mutex)
type chopstick struct {
	sync.Mutex
}

// Type philosopher, containing two chopsticks and the ID
type philosopher struct {
	id              int
	leftCS, rightCS *chopstick
}

// Philosopher's eat method
func (p philosopher) eat(c, end chan int) {

	for {
		// This loop will end once the philosopher receives a signal via end channel
		select {
		// If eat signal received, then perform the eat action locking and unlocking the chopsticks
		case nEats := <-c:
			p.leftCS.Lock()
			p.rightCS.Lock()

			nEats++
			fmt.Printf("Starting to eat %d\n", p.id)
			fmt.Printf("Finishing eating %d (%d times)\n", p.id, nEats)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			// Signal back that the eating was performed (with the counter value)
			c <- nEats
		case <-end:
			return
		}
	}
}

/*
The host function is the goroutine for the orchestrator.
It is in charge of handling the eating permissions for the philosophers
And he is in charge of coordinating that there are at most 2 philosophers eating at the same time
*/
func host(philos []*philosopher, wg *sync.WaitGroup) {

	// Channels to manage the eat acton (and times) per philosopher
	var nEatChannels [numEntities]chan int
	// Channels to end the philosophers' execution
	var abortChannels [numEntities]chan int

	// Counter to know the amount of times a philosopher was sent to eat
	var mealsPerPhilo [numEntities]int

	// Initialization of channels
	for i := range nEatChannels {
		nEatChannels[i] = make(chan int)
		abortChannels[i] = make(chan int)
	}

	// Set all the philosophers in the table
	for i := 0; i < numEntities; i++ {
		go philos[i].eat(nEatChannels[i], abortChannels[i])
	}

	// Initial philosopher to eat (and its partner, given that there must be 2 eating at once)
	// It should be a different one per execution
	initialPhilo := rand.Intn(numEntities)
	// And the partner
	secondPhilo := (initialPhilo + 2) % numEntities

	// Perform the eating process
	for continueEating(mealsPerPhilo) {
		// Signal the first philosopher to eat if has not eaten already all its amount
		if mealsPerPhilo[initialPhilo] < meals {
			nEatChannels[initialPhilo] <- mealsPerPhilo[initialPhilo]
		}
		// The same for its partner
		if mealsPerPhilo[secondPhilo] < meals {
			nEatChannels[secondPhilo] <- mealsPerPhilo[secondPhilo]
		}
		// Wait for the first philosopher to end to receive the updated amount of meals
		if mealsPerPhilo[initialPhilo] < meals {
			mealsPerPhilo[initialPhilo] = <-nEatChannels[initialPhilo]
		}
		// The same for the partner
		if mealsPerPhilo[secondPhilo] < meals {
			mealsPerPhilo[secondPhilo] = <-nEatChannels[secondPhilo]
		}

		// Increase indexes for philosophers
		initialPhilo = (initialPhilo + 1) % numEntities
		secondPhilo = (secondPhilo + 1) % numEntities
	}
	// Process done
	wg.Done()
}

// Function to determine when all the philosophers have eaten
func continueEating(mealsPerPhilo [numEntities]int) bool {
	for i := 0; i < numEntities; i++ {
		if mealsPerPhilo[i] < meals {
			return true
		}
	}
	return false
}

func main() {
	// Init the chopsticks
	csticks := make([]*chopstick, numEntities)
	for i := 0; i < numEntities; i++ {
		csticks[i] = new(chopstick)
	}

	// Init the philosophers
	philos := make([]*philosopher, numEntities)
	for i := 0; i < numEntities; i++ {
		philos[i] = &philosopher{i + 1, csticks[i], csticks[(i+1)%numEntities]}
	}

	// WaitGroup to prevent the host to be killed because of main goroutine termination
	var wg sync.WaitGroup

	// Add one element
	wg.Add(1)
	// Launch host goroutine with all the philosophers initialized
	go host(philos, &wg)
	// Wait for host to end
	wg.Wait()
	// Show termination message
	fmt.Println("All philosophers had lunch and the host finished its task.")
}
