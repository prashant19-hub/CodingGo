package main
import "fmt"

func hello(done chan bool) {
    fmt.Println("Hello from goroutine")
    done <- true
}

func main() {
    done := make(chan bool)
    go hello(done)
    <-done  // Blocks until received
    fmt.Println("Main done")
}
