package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	textMapInput, _ := os.ReadFile("./input_map.txt")
	WithInput(textMapInput)
	// textMap, _ := os.ReadFile("./test_map.txt")
	// WithInput(textMap)
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
	// visitedPlaces := map[string]string{}
	visitedPlaces := make([]string, 0)
	direction := "top"
	WardNextPlace := make([]int, 2)
	WardNextPlace[0] = wardCoordinates[0]
	WardNextPlace[1] = wardCoordinates[1]
	mapPlaces := make(map[string]bool)
	for WardInMap(wardCoordinates, mapLines) {
		maze[wardCoordinates[0]][wardCoordinates[1]] = "X"
		y := strconv.Itoa(wardCoordinates[1])
		x := strconv.Itoa(wardCoordinates[0])
		// visitedPlaces[x+y] = maze[wardCoordinates[0]][wardCoordinates[1]]
		visitedPlaces = appendIfUnique(visitedPlaces, x+":"+y)
		mapPlaces[x+":"+y] = true

		Move(WardNextPlace, direction)
		isValidWardInMap := WardInMap(WardNextPlace, mapLines)
		if !isValidWardInMap {
			fmt.Println("Ending in ", WardNextPlace)
			break
		}
		nextPlace := maze[WardNextPlace[0]][WardNextPlace[1]]

		if nextPlace == "#" {
			direction = DirectionChange(direction)

			WardNextPlace[0] = wardCoordinates[0]
			WardNextPlace[1] = wardCoordinates[1]

			Move(WardNextPlace, direction)
		}

		Move(wardCoordinates, direction)
		// printMap(maze, wardCoordinates, direction)
		y = strconv.Itoa(wardCoordinates[1])
		x = strconv.Itoa(wardCoordinates[0])
		// fmt.Printf("%s %s \n", x, y)
		// visitedPlaces[x+y] = maze[wardCoordinates[0]][wardCoordinates[1]]
		visitedPlaces = appendIfUnique(visitedPlaces, x+y)
	}

	printMap(maze, wardCoordinates, direction)
	fmt.Println("Total places", len(visitedPlaces), "ending", wardCoordinates, "Map", len(mapPlaces))
}
func printMap(maze [][]string, wardCoordinates []int, direction string) {
	buffer := "__Maze__\n"
	for x, char := range maze {
		for y, char2 := range char {
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
			} else {
				buffer += char2
			}
		}
		buffer += "\n"
	}
	buffer += "------\n"
	fmt.Println(buffer)
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
func WardInMap(wardCoordinates []int, mapLines []string) bool {
	return (wardCoordinates[0] < len(mapLines) && wardCoordinates[0] >= 0) && (wardCoordinates[1] < len(mapLines[0]) && wardCoordinates[1] >= 0)

}

func appendIfUnique(input []string, search string) []string {
	for x := 0; x < len(input); x++ {
		if input[x] == search {
			return input
		}
	}
	return append(input, search)
}
