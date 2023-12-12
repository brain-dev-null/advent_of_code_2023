package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type History []int

func isAllZero(values []int) bool {
    for _, value := range values {
	if value != 0 {
	    return false
	}
    }
    return true
}

func calculateDerivation(values []int) []int {
    derivation := []int{}
    for i := 0; i < len(values) - 1; i++ {
	currenct := values[i]
	next := values[i + 1]
	difference := next - currenct
	derivation = append(derivation, difference)
    }
    return derivation
}

func extrapolate(derivations [][]int) int {
    diff := 0
    for derivId := len(derivations) - 1; derivId >= 0; derivId-- {
	derivation := derivations[derivId]
	next := derivation[len(derivation) - 1] + diff
	diff = next
    }
    return diff
}

func extrapolateBackwards(derivations [][]int) int {
    diff := 0
    for derivId := len(derivations) - 1; derivId >= 0; derivId-- {
	derivation := derivations[derivId]
	next := derivation[0] - diff
	diff = next
    }
    return diff
}

func predictNextValue(history History) int {
    derivations := [][]int{history}
    for !isAllZero(derivations[len(derivations) - 1]){
	currentDerivation := derivations[len(derivations) - 1]
	nextDerivation := calculateDerivation(currentDerivation)
	derivations = append(derivations, nextDerivation)
    }
    predictedNextValue := extrapolate(derivations)
    return predictedNextValue
}

func predictPreviousValue(history History) int {
    derivations := [][]int{history}
    for !isAllZero(derivations[len(derivations) - 1]){
	currentDerivation := derivations[len(derivations) - 1]
	nextDerivation := calculateDerivation(currentDerivation)
	derivations = append(derivations, nextDerivation)
    }
    predictedNextValue := extrapolateBackwards(derivations)
    return predictedNextValue
}

func parseHistories(lines []string) []History {
    histories := []History{}
    for _, line := range lines {
	values := []int{}
	for _, rawNumber := range strings.Split(line, " ") {
	    number, err := strconv.Atoi(rawNumber)
	    if err != nil {
		log.Panicf("Failed to conver '%s' to number\n%d", rawNumber, err)
	    }
	    values = append(values, number)
	}
	histories = append(histories, values)
    }
    return histories
}

func partA(input []string) string {
    histories := parseHistories(input)
    result := 0
    for _, history := range histories {
	predictedValue := predictNextValue(history)
	result += predictedValue
    }
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    histories := parseHistories(input)
    result := 0
    for _, history := range histories {
	predictedValue := predictPreviousValue(history)
	result += predictedValue
    }
    return fmt.Sprintf("%d", result)
}

func main() {
    data, err := os.ReadFile("./data/day9.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

