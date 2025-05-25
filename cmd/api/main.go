package main

import (
	"fmt"
	"sync"
)

func main() {
    ch := make(chan int, 2)
    var wg sync.WaitGroup


    go func() {
		defer wg.Done()
		defer close(ch)
        fmt.Println("Goroutine: sending message")
		ch <- 1
		ch <- 2

        fmt.Println("Goroutine: message sent")
    }()


    // Read remaining values
	wg.Add(1)
	wg.Wait()
    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
    val, ok := <-ch
    fmt.Println(val, ok) // 0 false (0 is zero value of int)
}
