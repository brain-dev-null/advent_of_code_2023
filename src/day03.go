package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Location struct {
    x int
    y int
}

type PartNumber struct {
    number int
    locations []Location
}

type PartPositions [][]bool

type PartNumberPositions [][]int
type Gear Location
type GearRatio struct{a int; b int}

func addPartPosition(partPositions [][]bool, x int, y int) {
    adjacentLocations := []Location{
        {x-1, y-1}, {x, y-1}, {x+1, y-1},
        {x-1, y  }, {x, y  }, {x+1, y  },
        {x-1, y+1}, {x, y+1}, {x+1, y+1},
    }

    dimX := len(partPositions[0]) 
    dimY := len(partPositions)


    for _, location := range(adjacentLocations) {
        if location.x < 0 || dimX <= location.x {
            continue
        }
        if location.y < 0 || dimY <= location.y {
            continue
        }
        partPositions[location.y][location.x] = true
    }
}

func createNewPartNumber(rawNumber string, numberTailX int, numberY int) PartNumber {
    number, err := strconv.Atoi(rawNumber)
    if err != nil {
        log.Panicf("Error parsing strint '%s' to number!\n%d\n", rawNumber, err)
    }
    
    locations := []Location{}

    numberHeadX := numberTailX - (len(rawNumber) -1)

    for x := numberHeadX; x <= numberTailX; x++ {
        locations = append(locations, Location{x, numberY})
    }
    return PartNumber{number, locations}
}

func scanSchematicA(schematic []string) ([]PartNumber, PartPositions) {
    dimX := len(schematic[0])
    dimY := len(schematic)

    partPositions := make([][]bool, dimY)
    for y := range(partPositions) {
        partPositions[y] = make([]bool, dimX)
    }

    partNumbers := []PartNumber{}

    for y, line := range(schematic) {
        currentNumber := ""
        for x, char := range(line) {
            // Add the next digit to the current number
            if 48 <= char && char <= 57 {
                currentNumber += string(char)
            } else {
                // A part was found (char != '.') and its position is stored
                if char != 46 {
                    addPartPosition(partPositions, x, y)
                }
                // A number has ended and can now be stored
                if len(currentNumber) > 0 {
                    newPartNumber := createNewPartNumber(currentNumber, x-1, y)
                    currentNumber = ""
                    partNumbers = append(partNumbers, newPartNumber)
                }
            }
        }
        // A number has ended and can now be stored
        if len(currentNumber) > 0 {
            newPartNumber := createNewPartNumber(currentNumber, len(line) - 1, y)
            currentNumber = ""
            partNumbers = append(partNumbers, newPartNumber)
        }
    }

    return partNumbers, partPositions
}

func filterPartNumbers(partNumbers []PartNumber, partPositions PartPositions) []PartNumber {
    filteredPartNumbers := []PartNumber{}

    for _, partNumber := range(partNumbers) {
        for _, location := range(partNumber.locations) {
            if partPositions[location.y][location.x] {
                filteredPartNumbers = append(filteredPartNumbers, partNumber)
                break
            }
        }
    }

    return filteredPartNumbers
}

func addPartNumberPosition(index int, partNumber PartNumber, partNumberPositions PartNumberPositions) {
    for _, location := range(partNumber.locations) {
        // Store index with offset one -> 0 is used as sentinel value
        partNumberPositions[location.y][location.x] = index + 1
    }
}

func scanSchematicB(schematic []string) ([]PartNumber, PartNumberPositions, []Gear) {
    dimX := len(schematic[0])
    dimY := len(schematic)

    partNumberPositions := make([][]int, dimY)
    for y := range(partNumberPositions) {
        partNumberPositions[y] = make([]int, dimX)
    }

    partNumbers := []PartNumber{}
    gears := []Gear{}

    for y, line := range(schematic) {
        currentNumber := ""
        for x, char := range(line) {
            // Add the next digit to the current number
            if 48 <= char && char <= 57 {
                currentNumber += string(char)
            } else {
                // A gear was found (char == '.') and its position is stored
                if char == 42 {
                    gears = append(gears, Gear{x, y})
                }
                // A number has ended and can now be stored
                if len(currentNumber) > 0 {
                    newPartNumber := createNewPartNumber(currentNumber, x-1, y)
                    currentNumber = ""
                    partNumbers = append(partNumbers, newPartNumber)
                    addPartNumberPosition(len(partNumbers) - 1, newPartNumber, partNumberPositions)
                }
            }
        }
        // A number has ended and can now be stored
        if len(currentNumber) > 0 {
            newPartNumber := createNewPartNumber(currentNumber, len(line) - 1, y)
            currentNumber = ""
            partNumbers = append(partNumbers, newPartNumber)
            addPartNumberPosition(len(partNumbers) - 1, newPartNumber, partNumberPositions)
        }
    }

    return partNumbers, partNumberPositions, gears
}

func getGearRatios(partNumbers []PartNumber, partNumberPositions PartNumberPositions, gears []Gear) []GearRatio {
    dimY := len(partNumberPositions) 
    dimX := len(partNumberPositions[0])

    gearRatios := []GearRatio{}

    for _, gear := range(gears) {
        a := -1
        b := -1
        invalid := false
        startX := max(0, gear.x - 1)
        endX := min(dimX - 1, gear.x + 1)
        startY := max(0, gear.y - 1)
        endY := min(dimY - 1, gear.y + 1)

        for x := startX; x <= endX; x++ {
            for y:= startY; y <= endY; y++ {
                offsetIndex := partNumberPositions[y][x]

                // Sentinel value 0 found, no part number here
                if offsetIndex == 0 {
                    continue
                }

                index := offsetIndex - 1
                
                // Current part number already included
                if a == index || b == index {
                    continue
                }

                // New part number detected but already two stored
                if a != -1 && b != -1 {
                    invalid = false
                    break
                }

                // Set part number a first, then b 
                if a == -1 {
                    a = index
                } else {
                    b = index
                }
            }
            // Skipping this gear
            if invalid {
                break
            }
        }
        if !invalid && a != -1 && b != -1 {
            gearRatio := GearRatio{
                a: partNumbers[a].number,
                b: partNumbers[b].number,
            }
            gearRatios = append(gearRatios, gearRatio)
        }
    }
    return gearRatios
}

func partA(input []string) string {
    partNumbers, partPositions := scanSchematicA(input)

    relevantPartNumbers := filterPartNumbers(partNumbers, partPositions)
    
    result := 0
    for _, partNumber := range(relevantPartNumbers) {
        result += partNumber.number
    }
    return fmt.Sprintf("%d", result)
}

func partB(input []string) string {
    partNumbers, partNumberPositions, gears := scanSchematicB(input)
    gearRatios := getGearRatios(partNumbers, partNumberPositions, gears)
    
    result := 0
    for _, gearRatio := range(gearRatios) {
        result += gearRatio.a * gearRatio.b
    }
    return fmt.Sprintf("%d", result)
}

func main() {
    data, err := os.ReadFile("./data/day3.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

