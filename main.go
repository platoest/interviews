package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

// func main() {

// 	chain := blockchain.InitBlockChain()
// 	chain.AddBlock("First Block after Genesis")
// 	chain.AddBlock("Second Block after Genesis")
// 	chain.AddBlock("Third Block after Genesis")

// 	for _, block := range chain.Blocks {
// 		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
// 		fmt.Printf("Data in Block: %s\n", block.Data)
// 		fmt.Printf("Hash: %x\n", block.Hash)

// 		pow := blockchain.NewProof(block)
// 		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
// 		fmt.Println()
// 	}
//}

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("hello world")
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(wg)
	wg.Wait()
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
