package main

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type person struct {
	name string
}

func main() {
	runtime.GOMAXPROCS(6)

	now := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := http.Get("http://localhost:8080")
			if err != nil {
				log.Println(err)
			}
		}()
	}

	wg.Wait()
	log.Printf("since %fs\n", time.Since(now).Seconds())
}
