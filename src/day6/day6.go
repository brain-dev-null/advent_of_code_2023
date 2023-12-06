package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
    raceTime int
    recordDistance int
}

func midnightFormula(a float64, b float64, c float64) (float64, float64) {
    sqrt := math.Sqrt(b * b - 4 * a * c)
    x1 := (-b + sqrt) / 2 * a
    x2 := (-b - sqrt) / 2 * a
    return x1, x2
}

func getHoldTimeLimits(race Race) (int, int) {
    x1, x2 := midnightFormula(-1, float64(race.raceTime), -float64(race.recordDistance + 1))
    lowerLimit := int(math.Ceil(min(x1, x2)))
    upperLimit := int(math.Floor(max(x1, x2)))

    return lowerLimit, upperLimit
}

func parseValues(rawValues string, prefix string) []int {
    valuesString := strings.TrimPrefix(rawValues, prefix)
    values := []int{}
    for _, block := range(strings.Split(valuesString, " ")) {
        trimmed := strings.Trim(block, " ")
        if len(trimmed) == 0 {
            continue
        }
        value, err := strconv.Atoi(trimmed)
        if err != nil {
            log.Panicf("Failed to convet '%s' to int\n%d", trimmed, err)
        }
        values = append(values, value)
    }
    return values
}

func parseRaces(rawTimes string, rawDistances string) []Race {
    times := parseValues(rawTimes, "Time:")
    distances := parseValues(rawDistances, "Distance:")

    races := []Race{}
    for i, time := range(times) {
        distance := distances[i]
        race := Race{time, distance}
        races = append(races, race)
    }
    return races
}

func parseValueWithoutSpaces(rawValues string, prefix string) int {
    valuesString := strings.TrimPrefix(rawValues, prefix)
    digits := strings.ReplaceAll(valuesString, " ", "")
    value, err := strconv.Atoi(digits)
    if err != nil {
        log.Panicf("Failed to convet '%s' to int\n%d", digits, err)
    }
    return value
}

func parseSingleRace(rawTimes string, rawDistances string) Race {
    time := parseValueWithoutSpaces(rawTimes, "Time:")
    distance := parseValueWithoutSpaces(rawDistances, "Distance:")
    return Race{time, distance}
}

func partA(input []string) string {
    result := 1

    races := parseRaces(input[0], input[1])
    for _, race := range(races) {
        lowerLimit, upperLimit := getHoldTimeLimits(race)
        winningHoldTimesCount := upperLimit - lowerLimit + 1
        result *= winningHoldTimesCount
    }
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    race := parseSingleRace(input[0], input[1])
    lowerLimit, upperLimit := getHoldTimeLimits(race)
    result := upperLimit - lowerLimit + 1

    return fmt.Sprintf("%d", result)
}

func main() {
    data, err := os.ReadFile("./data/day6.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

