package main

import (
	"log"
	"os"
	"strings"
)

func partA(input []string) string {
    return ""
}

func partB(input []string) string {
    return ""
}

func main() {
    data, err := os.ReadFile("./data/day4.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

