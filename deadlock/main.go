package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof"
// 	"os"
// )

// func main() {
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()
// 	ch := make(chan int)

// 	s := make(chan os.Signal, 1)
// 	//signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

// 	select {
// 	case <-ch:
// 	case <-s:
// 		fmt.Println("stopped")
// 	}
// 	fmt.Println("Exit")
// }

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(m)
}
