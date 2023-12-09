package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
    code string
    left string
    right string
}
type Nodes map[string]Node

func step(instruction rune, currentNode Node, nodes Nodes) Node {
    if instruction == 'L' {
	return nodes[currentNode.left]
    } else {
	return nodes[currentNode.right]
    }
}

func followInstructions(startingNode Node, instructions string, nodes Nodes) int {
    steps := 0
    currentNode := startingNode
    for {
	for _, instruction := range(instructions) {
	    currentNode = step(instruction, currentNode, nodes)
	    steps++
	    if currentNode.code == "ZZZ" {
		return steps    	
	    }
	}
    }
}

func followInstructionsSingle(startingNode Node, instructions string, nodes Nodes) int {
    steps := 0
    currentNode := startingNode
    for {
	for _, instruction := range(instructions) {
	    currentNode = step(instruction, currentNode, nodes)
	    steps++
	    if strings.HasSuffix(currentNode.code, "Z") {
		return steps    	
	    }
	}
    }
}

func getStartingNodes(nodes Nodes) []Node {
    startingNodes := []Node{}
    for _, node := range(nodes) {
	if strings.HasSuffix(node.code, "A") {
	    startingNodes = append(startingNodes, node)
	}
    }
    return startingNodes
}

func allEqual(values []int) bool {
    equalValue := values[0]
    for _, value := range values {
	if value != equalValue {
	    return false
	}
    }
    return true
}

func gcd(a int, b int) int {
    for b != 0 {
	a, b = b, a % b
    }
    return a
}

func lcm(a int, b int, values ...int) int {
    result := a * b / gcd(a, b)

    for _, value := range values {
	result = lcm(result, value)
    }
    
    return result
}


func followParallelInstructions(instructions string, nodes Nodes) int {
    currentNodes := getStartingNodes(nodes)
    stepsToTarget := []int{}
    for _, currentNode := range currentNodes {
	steps := followInstructionsSingle(currentNode, instructions, nodes)
	stepsToTarget = append(stepsToTarget, steps)
    }
    steps := lcm(stepsToTarget[0], stepsToTarget[1], stepsToTarget[2:]...)
    return steps
}

func parseNode(line string ) Node {
    code, outbounds, _ := strings.Cut(line, " = ")
    outbounds = strings.Trim(outbounds, "()")
    leftCode, rightCode, _ := strings.Cut(outbounds, ", ")
    return Node{
	code: code,
	left: leftCode,
	right: rightCode,
    }
}

func parseMap(lines []string) (string, Nodes) {
    instructions := lines[0]
    nodes := map[string]Node{}
    for _, line := range(lines) {
	node := parseNode(line)
	nodes[node.code] = node
    }

    return instructions, nodes
}

func partA(input []string) string {
    instructions, nodes := parseMap(input)
    startNode := nodes["AAA"]
    stepsRequired := followInstructions(startNode, instructions, nodes)

    return fmt.Sprintf("%d", stepsRequired)
}

func partB(input []string) string {
    instructions, nodes := parseMap(input)
    stepsRequired := followParallelInstructions(instructions, nodes)

    return fmt.Sprintf("%d", stepsRequired)
}

func main() {
    data, err := os.ReadFile("./data/day8.txt")
    
    if err != nil {
        log.Fatal("%w", err)
    }

    text := strings.Split(string(data), "\n")

    resultA := partA(text[:len(text)-1])
    log.Printf("Result A: %s", resultA)

    resultB := partB(text[:len(text)-1])
    log.Printf("Result B: %s", resultB)
}

