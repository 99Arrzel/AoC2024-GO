package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Test")
	textMap, _ := os.ReadFile("./test_map.txt")
	WithInput(textMap)
	fmt.Println("First")
	textMapInput, _ := os.ReadFile("./input_map.txt")
	WithInput(textMapInput)
}
func WithInput(textMap []byte) {
	mapLines := strings.Split(string(textMap), "\n")
	maze := make([][]string, 0)
	wardCoordinates := make([]int, 2)
	wardFound := false
	for lineNumber, wordLine := range mapLines {
		splited := strings.Split(wordLine, "")
		maze = append(maze, splited)
		if !wardFound {
			for index, character := range splited {
				letterAsciiValue := int(character[0])
				if letterAsciiValue == 94 {
					wardCoordinates[0] = lineNumber
					wardCoordinates[1] = index
					wardFound = true
					break
				}
			}
		}
	}
	if !wardFound {
		fmt.Println("No ward")
		return
	}
	uniquePlaces, infiniteLoops, infinityLoopPlaces := evaluateUniquePlaces(wardCoordinates, maze)
	fmt.Println("Total places", len(uniquePlaces))
	fmt.Println("InfiniteLoops", infiniteLoops, len(infinityLoopPlaces))
}
func evaluateUniquePlaces(wardCoordinates []int, maze [][]string) (map[string]bool, int, map[string]bool) {
	direction := "top"
	mapPlaces := make(map[string]bool)
	infinityMapPlaces := make(map[string]bool)
	maxSize := len(maze)
	infiniteLoopsCounter := 0
	for WardInMap(wardCoordinates, maxSize) {
		mapPlaces[makeKey(wardCoordinates[0], wardCoordinates[1])] = true
		Move(wardCoordinates, direction)
		isValidWardInMap := WardInMap(wardCoordinates, len(maze))
		if !isValidWardInMap {
			fmt.Println("Ending in ", wardCoordinates)
			break
		}
		for maze[wardCoordinates[0]][wardCoordinates[1]] == "#" {
			StepBack(wardCoordinates, direction)
			direction = DirectionChange(direction)
			Move(wardCoordinates, direction)
		}
		nextMove := CopyMove(wardCoordinates, direction)
		if isMazeInfiniteLoopWith(maze, wardCoordinates, direction, nextMove) {
			infiniteLoopsCounter++
			infinityMapPlaces[makeKey(nextMove[0], nextMove[1])] = true
		}
	}
	return mapPlaces, infiniteLoopsCounter, infinityMapPlaces
}

// Evaluate if, with current maze, and ward position, this is a infinite loop
func isMazeInfiniteLoopWith(maze [][]string, startPosition []int, startDirection string, blockPosition []int) bool {
	position := copyPosition(startPosition)
	direction := string(startDirection)
	traversedPaths := make(map[string]bool)
	for WardInMap(position, len(maze)) {
		Move(position, direction)
		isValidWardInMap := WardInMap(position, len(maze))
		if !isValidWardInMap {
			return false
		}
		for maze[position[0]][position[1]] == "#" || (position[0] == blockPosition[0] && position[1] == blockPosition[1]) {
			StepBack(position, direction)
			direction = DirectionChange(direction)
			Move(position, direction)
			accumulator, changes := thisDirectionChanges(maze, position, direction, blockPosition)
			if traversedPaths[accumulator] {
				return true
			}
			if changes {
				traversedPaths[accumulator] = true
			}
		}
	}
	// fmt.Println("Bye", didTurn)
	return false
}
func thisDirectionChanges(maze [][]string, position []int, direction string, blockPosition []int) (string, bool) {
	newPos := copyPosition(position)
	accumulator := ""
	// printMap(maze, position, direction, blockPosition)
	for WardInMap(newPos, len(maze)) {
		accumulator += makeKey(newPos[0], newPos[1])
		if maze[newPos[0]][newPos[1]] == "#" || (newPos[0] == blockPosition[0] && newPos[1] == blockPosition[1]) {
			return accumulator, true
		}
		Move(newPos, direction)
	}
	return accumulator, false
}
func copyPosition(position []int) []int {
	newPos := make([]int, 2)
	newPos[0] = position[0]
	newPos[1] = position[1]
	return newPos
}
func Move(wardCoordinates []int, direction string) {
	if direction == "top" {
		wardCoordinates[0]--
	} else if direction == "right" {
		wardCoordinates[1]++
	} else if direction == "bottom" {
		wardCoordinates[0]++
	} else if direction == "left" {
		wardCoordinates[1]--
	}
}
func CopyMove(wardCoordinates []int, direction string) []int {
	newPos := make([]int, 2)
	newPos[0] = wardCoordinates[0]
	newPos[1] = wardCoordinates[1]
	if direction == "top" {
		newPos[0]--
	} else if direction == "right" {
		newPos[1]++
	} else if direction == "bottom" {
		newPos[0]++
	} else if direction == "left" {
		newPos[1]--
	}
	return newPos
}
func StepBack(wardCoordinates []int, direction string) {
	if direction == "top" {
		wardCoordinates[0]++
	} else if direction == "right" {
		wardCoordinates[1]--
	} else if direction == "bottom" {
		wardCoordinates[0]--
	} else if direction == "left" {
		wardCoordinates[1]++
	}
}
func DirectionChange(direction string) string {
	if direction == "top" {
		return "right"
	}
	if direction == "right" {
		return "bottom"
	}
	if direction == "bottom" {
		return "left"
	}
	return "top"
}

// Maze is square
func WardInMap(wardCoordinates []int, maxSize int) bool {
	return (wardCoordinates[0] < maxSize && wardCoordinates[0] >= 0) && (wardCoordinates[1] < maxSize && wardCoordinates[1] >= 0)
}
func makeKey(a int, b int) string {
	x := strconv.Itoa(a)
	y := strconv.Itoa(b)
	return "[" + x + ":" + y + "]"
}
func printMap(maze [][]string, wardCoordinates []int, direction string, obstacle []int) {
	buffer := "__Maze__\n"
	for x, char := range maze {
		for y, char2 := range char {
			if obstacle[0] == x && obstacle[1] == y {
				buffer += "O"
				continue
			}
			if wardCoordinates[0] == x && wardCoordinates[1] == y {
				if direction == "top" {
					buffer += "^"
				}
				if direction == "right" {
					buffer += ">"
				}
				if direction == "bottom" {
					buffer += "v"
				}
				if direction == "left" {
					buffer += "<"
				}
				continue
			}
			if char2 == "^" {
				buffer += "."
				continue
			}
			buffer += char2
		}
		buffer += "\n"
	}
	buffer += "------\n"
	fmt.Println(buffer)
}
