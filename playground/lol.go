package main

import (
	"context"
	"fmt"
	"time"
)

func readWord() (word string) {
	fmt.Println("Type a word, then hit enter")
	fmt.Scanf("%s", &word)
	return
}

func main() {
	ch := make(chan string)
	go func() { ch <- readWord() }()
	ctx, f := context.WithTimeout(context.Background(), 5*time.Second)

	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-ctx.Done():
		fmt.Println("Timeout.	")
	}
}
