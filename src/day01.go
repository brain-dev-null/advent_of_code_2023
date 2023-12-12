package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getCalibrationValue(line string) int {
    first := 0
    last := 0

    for _, char := range line {
        if 48 <= char && char <= 57 {
            first = int(char) - 48
            break
        }
    }

    for i := len(line) - 1; i >= 0; i-- {
        char := line[i]
        if 48 <= char && char <= 57 {
            last = int(char) - 48
            break
        }
        
    }

    return 10 * first + last
}

func getTextualNumber(text string) int {
    if len(text) >= 3 {
        if text[:3] == "one" {
            return 1
        } else if text[:3] == "two" {
            return 2
        } else if text[:3] == "six" {
            return 6
        }
    }

    if len(text) >= 4 {
         if text[:4] == "zero" {
            return 0
        } else if text[:4] == "four" {
            return 4
        } else if text[:4] == "five" {
            return 5
        } else if text[:4] == "nine" {
            return 9
        }
           
    }

    if len(text) >= 5 {
        if text[:5] == "three" {
            return 3
        } else if text[:5] == "seven" {
            return 7
        } else if text[:5] == "eight" {
            return 8
        }
    }

    return -1
}

func getCalibrationValueTextual(line string) int {
    nums := []int{}
    for i := 0; i < len(line); i++ {
        char := line[i]
        if 48 <= char && char <= 57 {
            nums = append(nums, int(char) - 48)
            continue
        } 

        textualNumber := getTextualNumber(line[i:])
        if textualNumber > -1 {
            nums = append(nums, textualNumber)
        }
    }
    first := nums[0]
    last := nums[len(nums) - 1]

    return first * 10 + last
}

func partA(input []string) string {
    calValSum := 0
    for _, line := range(input) {
        calValSum += getCalibrationValue(line)
    }
    return fmt.Sprintf("%d", calValSum)
}

func partB(input []string) string {
    calValSum := 0
    for _, line := range(input) {
        calValSum += getCalibrationValueTextual(line)
    }
    return fmt.Sprintf("%d", calValSum)
}

func main() {
    data, err := os.ReadFile("./data/day1.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text)
    resultB := partB(text[:len(text)-1])

    log.Printf("Result A: %s", resultA)
    log.Printf("Result B: %s", resultB)
}

