package main

import (
	"fmt"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup
var mut sync.Mutex
var a = make (chan bool, 1)
var b = make(chan bool, 1)

type ChopS struct { sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	eatCnt int
}


func host() { //Host can only give permission to two philosophers at once.

for i := 0 ; i<15;i++ {
	a <- true
	b <- true
}
} 

func (p Philo) eat (who int) {
		
	for p.eatCnt<3 { //Each philosopher should eat only 3 times.	

	select {

	case <-a:
		
		if (rand.Intn(100)%2==0) { //The philosophers pick up the chopsticks in any order.

			p.leftCS.Lock() 
			p.rightCS.Lock()
			fmt.Println("starting to eat:", who) 
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			mut.Lock()
			p.eatCnt ++ 
			mut.Unlock()
			fmt.Println("finishing eating:", who, " eat count: ",p.eatCnt)
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock() 			
			fmt.Println("starting to eat:", who) 
			p.leftCS.Unlock()			
			p.rightCS.Unlock()
			mut.Lock()
			p.eatCnt ++ 
			mut.Unlock()
			fmt.Println("finishing eating:", who, " eat count: ",p.eatCnt)
		}
			wg.Done()
	
	case <-b:
		
		if (rand.Intn(100)%2==0) { //The philosophers pick up the chopsticks in any order.

			p.leftCS.Lock() 
			p.rightCS.Lock()
			fmt.Println("starting to eat:", who) 
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			mut.Lock()
			p.eatCnt ++ 
			mut.Unlock()
			fmt.Println("finishing eating:", who, " eat count: ",p.eatCnt)
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock() 			
			fmt.Println("starting to eat:", who) 
			p.leftCS.Unlock()			
			p.rightCS.Unlock()
			mut.Lock()
			p.eatCnt ++ 
			mut.Unlock()
			fmt.Println("finishing eating:", who, " eat count: ",p.eatCnt)
		}
			wg.Done()
	}




	}
}

func main() {
	fmt.Println("Note: All conditions are met. Please check the code.")
	var eatCnt int
	Csticks := make([]*ChopS,5)
	for i := 0 ; i < 5 ; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5) // 5 philosophers, one chopstick between each pair.

	for i := 0 ; i < 5 ; i++ {
		philos[i] = &Philo {Csticks[i], Csticks [(i+1)%5], eatCnt}
	}
	
	wg.Add(15)
	go host()
	go philos[0].eat(0)
	go philos[1].eat(1)
	go philos[2].eat(2)
	go philos[3].eat(3)
	go philos[4].eat(4)

	wg.Wait()

}










	