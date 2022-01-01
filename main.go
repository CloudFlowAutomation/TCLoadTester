package main

import (
	"net/http"
	"sync"
)


func main() {
	var wg sync.WaitGroup

	for i := 1; 1 <= 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ := http.Get("https://darkmatterit.io")
		}()
	}
	wg.Wait()
}