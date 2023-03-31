package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan string)
	go func() {
		count("essi", c)
		count("moeen", c)
		wg.Done()
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		fmt.Println("shit", thing)
		time.Sleep(time.Millisecond * 100)
	}
}
