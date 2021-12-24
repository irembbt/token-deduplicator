package generate

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"time"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyz")

const token_length = 7

func randTokens(n int) []byte {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return b
}

func generateTokens(size int) []byte {
	items := make([][]byte, size)

	for i := 0; i < size; i++ {
		items[i] = randTokens(token_length)
	}

	return bytes.Join(items, []byte("\n"))
}

func WriteRandomTokens(filename string, size int) {
	rand.Seed(time.Now().UnixNano())

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Delete file content
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	//Add file content
	fileContent := generateTokens(size)

	f.Write(fileContent)
}
