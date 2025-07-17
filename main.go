package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"sync"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
		log.Fatal(runHTTPServer())
	}()
	go func() {
		defer wg.Done()
		log.Fatal(runTCPServer())
	}()

	wg.Wait()
}
