package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
    winningNumbers map[int]struct{}
    cardNumbers []int
}

func pow(base int, exponent int) int {
    if exponent == 0 {
        return 1
    }

    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }

    return result
}

func parseCard(rawCard string) Card {
    _, numbers, _ := strings.Cut(rawCard, ": ")
    
    rawWinningNumbers, rawCardNumbers, _ := strings.Cut(numbers, "|")
    rawWinningNumbers = strings.Trim(rawWinningNumbers, " ")
    rawCardNumbers = strings.Trim(rawCardNumbers, " ")
    
    winningNumbers := map[int]struct{}{}
    for _, rawWinningNumber := range(strings.Split(rawWinningNumbers, " ")) {
        if len(rawWinningNumber) == 0 {
            continue
        }
        winningNumber, err := strconv.Atoi(rawWinningNumber)
        if err != nil {
            log.Panicf("Error parsing string '%s' as int:\n%d", rawWinningNumber, err)
        }
        winningNumbers[winningNumber] = struct{}{}
    }
    
    cardNumbers := []int{}
    for _, rawCardNumber := range(strings.Split(rawCardNumbers, " ")) {
        if len(rawCardNumber) == 0 {
            continue
        }
        cardNumber, err := strconv.Atoi(rawCardNumber)
        if err != nil {
            log.Panicf("Error parsing string '%s' as int:\n%d", rawCardNumber, err)
        }
        cardNumbers = append(cardNumbers, cardNumber)
    }


    return Card{winningNumbers, cardNumbers}
}


func parseCards(rawCards []string) []Card {
    cards := []Card{}
    for _, rawCard := range(rawCards) {
        card := parseCard(rawCard)
        cards = append(cards, card)
    }
    return cards
}

func getWinCount(card Card) int {
    counter := 0
    for _, cardNumber := range(card.cardNumbers) {
        _, isWinning := card.winningNumbers[cardNumber]
        if isWinning {
            counter += 1
        }
    }
    return counter
}


func partA(input []string) string {
    result := 0
    cards := parseCards(input)
    for _, card := range(cards) {
        winCount := getWinCount(card)
        if winCount > 0 {
            result += pow(2, winCount - 1)
        }
    }
    
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    result := 0
    cards := parseCards(input)
    additionalCopies := map[int]int{}

    for cardNumber, card := range(cards) {
        // Number of winning numbers on this card
        winCount := getWinCount(card)

        // Total number of copies for this card
        cardCount := 1 + additionalCopies[cardNumber]

        // Add this card's copies to the total count
        result += cardCount
        
        // For each copy of this card, add one for each of the next 'winCount' cards
        for i := 1; i <= winCount; i++ {
            additionalCopies[cardNumber + i] += cardCount
        }
    }
    
    return fmt.Sprintf("%d", result)

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

