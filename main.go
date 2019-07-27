package main

import (
	"fmt"
	"sync"
        "sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go increment("Proc01:")
	go increment(" Proc02:")
	wg.Wait()
        fmt.Println("Final Count:", counter)
}

func increment(s string) {
	for i := 0; i < 50; i++ {
                atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Count:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
