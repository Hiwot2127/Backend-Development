package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// WordFrequencyCounter takes a string as input and returns a map with the frequency of each word
func wordFrequency(text string) map[string]int {
    re := regexp.MustCompile(`[^\w\s]`)
    newtext := re.ReplaceAllString(text, "")
    newtext = strings.ToLower(newtext)
    words := strings.Fields(newtext)

    // a map to keep track of word frequencies
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    return wordCount
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    validInput := false
    var text string

    // Regular expression to validate input contains only letters and spaces
    validStringRe := regexp.MustCompile(`^[a-zA-Z\s]+$`)

    for !validInput {
        fmt.Print("Enter a string: ")
        text, _ = reader.ReadString('\n')
        text = strings.TrimSpace(text)
        
        if len(text) == 0 {
            fmt.Println("Empty input. Please enter a valid string.")
            continue
        }

        if !validStringRe.MatchString(text) {
            fmt.Println("Invalid input. Please enter a string containing only letters and spaces.")
            continue
        }

        validInput = true
    }
  
    //Printing the result
    wordFreq := wordFrequency(text)
    fmt.Println("Word Frequency Count:")
    for word, count := range wordFreq {
        fmt.Printf("%s: %d\n", word, count)
    }
}
