package dedupe

//HashDeduplicate deduplicates a slice of tokens by putting them in a hash table.
//Every token that didn't already exist in hash table are published to the unique tokens channel.
//Hash table also stores frequencies for all tokens, tokens that occur only once are filtered
//after all tokens are read.
//Duplicate tokens and their occurence frequencies are returned as a map.
func HashDeduplicate(tokens []string, unique_channel chan string) map[string]int {
	token_map := make(map[string]int, 10000000)
	duplicate_tokens := make(map[string]int)

	for _, token := range tokens {
		_, ok := token_map[token]
		if ok {
			token_map[token]++
		} else {
			token_map[token] = 1
			unique_channel <- token
		}
	}
	close(unique_channel)
	for token, frequency := range token_map {
		if frequency > 1 {
			duplicate_tokens[token] = frequency
		}
	}
	return duplicate_tokens
}
