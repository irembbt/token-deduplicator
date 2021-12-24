package main

import (
	"d4l_token/database"
	"d4l_token/dedupe"
	"d4l_token/generate"
	"d4l_token/read"
	"log"
	"sync"
)

func main() {
	const filename = "tokens.txt"
	generate.WriteRandomTokens(filename, 10000000)

	tokens := read.GetTokens(filename)
	unique_chan := make(chan string, 10000)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go readyDb(unique_chan, wg)

	duplicate_tokens := dedupe.SortDeduplicate(tokens, unique_chan)
	// duplicate_tokens := dedupe.HashDeduplicate(tokens, unique_chan)

	for key, val := range duplicate_tokens {
		log.Printf("Token: %s, Frequency: %d\n", key, val)
	}
	log.Printf("Duplicate tokens count: %d\n", len(duplicate_tokens))
	wg.Wait()
}

func readyDb(token_chan chan string, wg *sync.WaitGroup) {
	database.InitDb()
	database.Migrate()
	database.SaveTokens(token_chan)
	wg.Done()
}
