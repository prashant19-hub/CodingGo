package main

import (
    "fmt"
    "sync"
)

func generate(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < n; i++ {
            ch <- i
        }
    }()
    return ch
}

func multiply(in <-chan int, factor int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * factor
        }
    }()
    return out
}

func filter(in <-chan int, pred func(int) bool) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if pred(n) {
                out <- n
            }
        }
    }()
    return out
}

func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer wg.Done()
        nums := generate(20)
        doubled := multiply(nums, 2)
        evens := filter(doubled, func(n int) bool { return n%2 == 0 })
        for n := range evens {
            fmt.Println(n)  // 0,4,8,12,16,20,24,28,32,36
        }
    }()
    wg.Wait()
}

