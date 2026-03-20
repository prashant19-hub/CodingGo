package main

import (
	"fmt"
	"sync"
)

func main(){
	mutex := sync.Mutex{}
	accountinbalance := 0
	wg := sync.WaitGroup{}
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go balancegrow(&accountinbalance, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(accountinbalance)

}

func balancegrow(balance *int ,wg *sync.WaitGroup , lock *sync.Mutex){
	defer wg.Done()
	lock.Lock()             //read to write , then write time  the lock  using and work
	*balance++
	lock.Unlock()          //unclock they after read

}