package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)
type Hand [5]rune
type Player struct {
    hand Hand
    handValue int
    bid int
}

func (hand Hand) printHand() string {
    str := ""
    for _, c := range(hand) {
        str += string(c)
    }
    return str
}

func getHandKind(hand Hand) int {
    groups := map[rune]int{}
    for _, card := range(hand) {
        groups[card] += 1
    }

    pairs := 0
    trio := false

    
    for _, groupSize := range(groups) {
        if groupSize == 5 {
            return 7
        }
        if groupSize == 4 {
            return 6
        }
        if groupSize == 3 {
            trio = true
        }
        if groupSize == 2 {
            pairs += 1
        }
    }

    if trio && pairs == 1 {
        return 5
    }

    if trio {
        return 4
    }

    if pairs == 2 {
        return 3
    }

    if pairs == 1 {
        return 2
    }

    return 1
}

func getHandKindJokers(hand Hand) int {
    groups := map[rune]int{}
    jokers := 0
    for _, card := range(hand) {
        if card == 'J'{
            jokers += 1
        } else {
            groups[card] += 1
        }
    }

    pairs := 0
    trio := false
    four := false
    five := false
    
    for _, groupSize := range(groups) {
        if groupSize == 2 {
            pairs += 1
        } else if groupSize == 3 {
            trio = true
        } else if groupSize == 4 {
            four = true
        } else if groupSize == 5 {
            five = true
        }

    }


    // Five of a kind
    if jokers == 5 || jokers == 4 || (jokers == 3 && pairs == 1) || (jokers == 2 && trio) || (jokers == 1 && four) || five {
        return 7
    }

    // Four of a kind
    if jokers == 3 || (jokers == 2 && pairs == 1) || (jokers == 1 && trio) || four {
        return 6
    }

    // Full house
    if (jokers == 2 && trio) || (jokers == 1 && pairs == 2) ||(pairs == 1 && trio) {
        return 5
    }

    // Three of a kind
    if jokers == 2 || (jokers == 1 && pairs == 1) || trio {
        return 4
    }

    // Two pair 
    if (jokers == 1 && pairs ==1) || pairs == 2 {
        return 3
    }

    // One pair
    if jokers == 1 || pairs == 1 {
        return 2
    }

    return 1
}

func getCardValue(card rune) int {
    if '2' <= card && card <= '9' {
        return int(card - '2')
    }
    return map[rune]int{
        'T': 10,
        'J': 11,
        'Q': 12,
        'K': 13,
        'A': 14,
    }[card]
}

func getCardValueJokers(card rune) int {
    if '2' <= card && card <= '9' {
        return int(card - '0')
    }
    return map[rune]int{
        'J':  1,
        'T': 10,
        'Q': 12,
        'K': 13,
        'A': 14,
    }[card]
}

func getHandValue(hand Hand, getKind func(Hand) int, getValue func(rune) int) int {
    handValue := getKind(hand)
    for _, card := range(hand) {
        handValue *= 100
        handValue += getValue(card)
    }
    return handValue
}

// Reverse sort the players
func sortPlayers(players []Player) {
    slices.SortStableFunc(players, func(a, b Player) int {
        return cmp.Compare(a.handValue, b.handValue)
    })
}

func parsePlayer(line string, getKind func(Hand) int, getValue func(rune) int) Player {
    hand := [5]rune{}
    for i, card := range(line[:5]) {
        hand[i] = card
    }
    _, rawBid, _ := strings.Cut(line, " ")
    bid, err := strconv.Atoi(rawBid)
    if err != nil {
        log.Panicf("Failed to convert '%s' to int", rawBid)
    }
    handValue := getHandValue(hand, getKind, getValue)
    return Player{
        hand: hand,
        handValue: handValue,
        bid: bid,
    }
}

func parsePlayers(lines []string, getKind func(Hand) int, getValue func(rune) int) []Player {
    players := []Player{}
    for _, line := range(lines) {
        player := parsePlayer(line, getKind, getValue)
        players = append(players, player)
    }
    return players
}


func partA(input []string) string {
    players := parsePlayers(input, getHandKind, getCardValue)
    sortPlayers(players)
    result := 0
    for i, player := range(players) {
        rank := i + 1
        result += rank * player.bid
    }
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    players := parsePlayers(input, getHandKindJokers, getCardValueJokers)
    sortPlayers(players)
    result := 0
    for i, player := range(players) {
        rank := i + 1
        result += rank * player.bid
    }
    return fmt.Sprintf("%d", result)
}

func main() {
    data, err := os.ReadFile("./data/day7.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    log.Printf("Result A: %s", resultA)
    resultB := partB(text[:len(text)-1])

    log.Printf("Result B: %s", resultB)
}

