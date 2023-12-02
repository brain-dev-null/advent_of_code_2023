package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reveal struct {
    reds int
    greens int
    blues int
}

type Constraint struct {
    reds int
    greens int
    blues int
}

type Game []Reveal

func parseAndAddColor(colorCount string, reveal *Reveal) {
    rawCount, color, _ := strings.Cut(colorCount, " ")
    count, err := strconv.Atoi(rawCount)

    if err != nil {
        log.Panicf("%d", err)
    }

    if color == "red" {
        reveal.reds = count
    } else if color == "green" {
        reveal.greens = count
    } else if color == "blue" {
        reveal.blues = count
    }
}


func parseRawReveal(rawReveal string) Reveal {
    reveal := Reveal{}
    for _, colorCount := range(strings.Split(rawReveal, ", ")) {
        parseAndAddColor(colorCount, &reveal)
    }
    return reveal
}

func parseGame(line string) Game {
    game := []Reveal{}
    _, revealsText, _ := strings.Cut(line, ": ")
    rawReveals := strings.Split(revealsText, "; ")

    for _, rawReveal := range(rawReveals) {
        reveal := parseRawReveal(rawReveal)
        game = append(game, reveal)
    }

    return game
}

func testConstraint(game Game, constraint Constraint) bool {
    maxRed := 0
    maxGreen := 0
    maxBlue := 0

    for _, reveal := range(game) {
        maxRed = max(reveal.reds, maxRed)
        maxGreen = max(reveal.greens, maxGreen)
        maxBlue = max(reveal.blues, maxBlue)
    }

    redTooHigh := maxRed > constraint.reds
    greenTooHigh := maxGreen > constraint.greens
    blueTooHigh := maxBlue > constraint.blues

    return !(redTooHigh || greenTooHigh || blueTooHigh)
}

func generateConstraint(game Game) Constraint {
    maxRed := 0
    maxGreen := 0
    maxBlue := 0

    for _, reveal := range(game) {
        maxRed = max(reveal.reds, maxRed)
        maxGreen = max(reveal.greens, maxGreen)
        maxBlue = max(reveal.blues, maxBlue)
    }
    
    return Constraint{reds: maxRed, greens: maxGreen, blues: maxBlue}
}

func partA(input []string) string {
    constraint := Constraint{reds: 12, greens: 13, blues: 14}
    result := 0
    for gameId, line := range input {
        game := parseGame(line)
        valid := testConstraint(game, constraint)
        if valid {
            result += gameId + 1
        }
    }
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    result := 0
    for _, line := range input {
        game := parseGame(line)
        constraint := generateConstraint(game)
        result += constraint.reds * constraint.greens * constraint.blues
    }
    return fmt.Sprintf("%d", result)
}

func main() {
    data, err := os.ReadFile("./data/day2.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

