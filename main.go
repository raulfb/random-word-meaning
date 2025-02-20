package main

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var words []string
var rng *rand.Rand // Random number generator

func loadWords() error {
	file, err := os.Open("words.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return scanner.Err()
}

// Get a random word
func getRandomWord() string {
	return words[rng.Intn(len(words))]
}

// Get the meaning of a word
func getMeaning(word string) (string, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", nil // Meaning not found
	}

	var result []struct {
		Meanings []struct {
			Definitions []struct {
				Meaning string `json:"definition"`
			} `json:"definitions"`
		} `json:"meanings"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result) > 0 && len(result[0].Meanings) > 0 && len(result[0].Meanings[0].Definitions) > 0 {
		return result[0].Meanings[0].Definitions[0].Meaning, nil
	}

	return "", nil // Meaning not found
}

func main() {
	// Load words at startup
	err := loadWords()
	if err != nil {
		panic(err)
	}

	// Initialize the random number generator
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil) // Render the template
	})

	r.GET("/random-word", func(c *gin.Context) {
		word := getRandomWord()
		meaning, err := getMeaning(word)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching meaning"})
			return
		}

		response := map[string]string{
			"word":    word,
			"meaning": meaning,
		}
		c.JSON(http.StatusOK, response) // Return the word and its meaning in JSON format
	})

	r.Run(":8081") // Start the server on port 8081
}
