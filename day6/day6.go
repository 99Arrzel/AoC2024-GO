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
	textMap, _ := os.ReadFile("./test_map.txt")
	WithInput(textMap)
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
	}
	WardInMapCondition := WardInMap(wardCoordinates, mapLines)
	visitedPlaces := map[string]string{}
	direction := "top"
	WardNextPlace := make([]int, 2)
	WardNextPlace[0] = wardCoordinates[0]
	WardNextPlace[1] = wardCoordinates[1]
	counter := 0
	for WardInMapCondition {
		y := strconv.Itoa(wardCoordinates[1])
		x := strconv.Itoa(wardCoordinates[0])
		visitedPlaces[x+y] = maze[wardCoordinates[0]][wardCoordinates[1]]
		counter++
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
		WardInMapCondition = WardInMap(wardCoordinates, mapLines)
		// printMap(maze, wardCoordinates, direction)
		y = strconv.Itoa(wardCoordinates[1])
		x = strconv.Itoa(wardCoordinates[0])
		visitedPlaces[x+y] = maze[wardCoordinates[0]][wardCoordinates[1]]
	}
	fmt.Println("Total places", len(visitedPlaces), "counter", counter)
}
func printMap(maze [][]string, wardCoordinates []int, direction string) {
	fmt.Println("__Maze__")
	for x, char := range maze {
		for y, char2 := range char {
			if wardCoordinates[0] == x && wardCoordinates[1] == y {
				if direction == "top" {
					fmt.Print("^")
				}
				if direction == "right" {
					fmt.Print(">")
				}
				if direction == "bottom" {
					fmt.Print("v")
				}
				if direction == "left" {
					fmt.Print("<")
				}
			} else {
				fmt.Print(char2)
			}
		}
		fmt.Println()
	}
	fmt.Println("________")
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
