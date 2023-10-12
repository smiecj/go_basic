package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const (
	routineCount = 2000
)

func main() {
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	for i := 0; i < routineCount; i++ {
		go func() {
			time.Sleep(100 * time.Minute)
		}()
	}

	fmt.Printf("%d routines have created, main ready to sleep", routineCount)
	time.Sleep(100 * time.Minute)
}
