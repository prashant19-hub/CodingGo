package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Typing time test")
	fmt.Println("Press ENTER to start...")
	_, _ = reader.ReadString('\n')

	start := time.Now()

	fmt.Println("Type a sentence and press ENTER:")
	_, _ = reader.ReadString('\n')

	elapsed := time.Since(start)

	fmt.Printf("You took %s to type your sentence.\n", elapsed)
}
