package main

import (
	"fmt"
	"time"
)

func main() {

	// 1st type of GoRoutines
	//parallel
	// -----------|----------|-----|-----> process a or function a
	// -----------|----------|-----|-----> process b or function b
	// -----------|----------|-----|-----> process c or function c

	//cores (cores are running in parallel method)
	// |1-> main|   |2 -> a|
	// |3 -> b|     |4 -> c|

	//  2nd types of GoRoutines
	// Concurrency
	// ------ -----   | -----|-----  -  - ------>process d or funtion d
	//       -     --- -           -- -- -      >processe or funtion e

	//cores(core are running concurrency method)
	// |1-> main or d|   |2 -> a or d,e|
	// |3 -> b|          |4 -> c|
    

	   //parallel / linear
// a()
// b()
// c()
// d()

                                    //concurrency method but they are run and not run because there in main func mm sec in closed but time sec in fixed time and there time in maximum to maximum time extra run 
go a()
go b()
go c()
go d()

time.Sleep(time.Second * 3)        //concurrency methods worked they are use to hold the youe timing by main func delay run ,means time.Sleep(time.Second * "time sec/min in")

}

func a() {
	for i := 1; i<= 3; i++{
     fmt.Println("Func a run hua , i =",i)
	}
}
func b() {
	for i := 1; i<= 3; i++{
     fmt.Println("Func b run hua , i =",i)
	}
}
func c() {
	for i := 1; i<= 3; i++{
	fmt.Println("Func c run hua, i =",i)
}
}
func d() {
	for i := 1; i<= 3; i++{
	fmt.Println("Func d run hua , i =",i)
}
}

