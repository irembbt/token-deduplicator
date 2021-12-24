package read

import (
	"bufio"
	"log"
	"os"
)

func GetTokens(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	tokens := make([]string, 0, 10000000)

	//Read file contents line by line
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	return tokens
}
