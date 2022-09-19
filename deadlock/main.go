package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	ch := make(chan int)

	s := make(chan os.Signal, 1)
	//signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ch:
	case <-s:
		fmt.Println("stopped")
	}
	fmt.Println("Exit")
}
