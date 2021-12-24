package dedupe

import (
	"sort"
)

//SortDeduplicate deduplicates a slice of tokens by sorting them in place
//and iterating over the sorted tokens to detect duplicates.
//Duplicate tokens will be consecutive elements in the sorted slice
//therefore we can check if a token is duplicate by comparing to the previous token.
//Duplicate tokens and their occurence frequencies are returned as a map
//and deduplicate tokens are published to the channel
func SortDeduplicate(tokens []string, unique_channel chan string) map[string]int {

	sort.Strings(tokens)

	prev_token := ""
	duplicate_tokens := make(map[string]int)
	for _, token := range tokens {
		if prev_token == token {
			_, ok := duplicate_tokens[token]
			if !ok {
				duplicate_tokens[token] = 1
			}
			duplicate_tokens[token]++
		} else {
			unique_channel <- token
		}
		prev_token = token
	}
	close(unique_channel)
	return duplicate_tokens
}
