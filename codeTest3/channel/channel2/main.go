package main

import (
    "context"
    "fmt"
)

func integers(ctx context.Context) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; ; i++ {
            select {
            case <-ctx.Done():
                return
            case ch <- i:
            }
        }
    }()
    return ch
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3)
    defer cancel()

    ints := integers(ctx)
    for n := range ints {
        fmt.Println(n)  // Prints 0,1,2... until timeout
    }
    fmt.Println("Done")
}
