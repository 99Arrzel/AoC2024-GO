package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseWord(word string) string {
	newWord := ""
	for x := len(word) - 1; x >= 0; x-- {
		newWord = newWord + string(word[x])
	}
	return newWord
}
func rotateArray(matrixs [][]string) [][]string {
	width := len(matrixs)
	height := len(matrixs[0])
	newMatrixs := make([][]string, 0) //Assuming it is squarish
	for x := 0; x < height; x++ {
		newRow := make([]string, width)
		newMatrixs = append(newMatrixs, newRow)
	}
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			newMatrixs[x][y] = matrixs[y][x]
		}
		newMatrixs[x] = strings.Split(reverseWord(strings.Join(newMatrixs[x], "")), "")
	}
	return newMatrixs
}
func inputToMatrixs(input string) [][]string {
	newMatrixs := make([][]string, 0)
	rowValues := strings.Split(input, "\n")
	for x := 0; x < len(rowValues); x++ {
		lineValue := strings.Split(rowValues[x], "")
		newMatrixs = append(newMatrixs, lineValue)
	}
	return newMatrixs
}

func inlineSearch(input [][]string, search string) int {
	total := 0
	for x := 0; x < len(input); x++ {
		total += strings.Count(strings.Join(input[x], ""), search)
	}
	return total
}
func countInArray(input []string, search string) int {
	total := 0
	for x := 0; x < len(input); x++ {
		total += strings.Count(input[x], search)
	}
	return total
}
func evaluateMAS(input string) bool {
	validSequences := []string{"MSMS", "SMMS", "SMSM", "MSSM"}
	for x := 0; x < len(validSequences); x++ {
		if strings.Contains(input, validSequences[x]) {
			return true
		}
	}
	return false
}
func appendIfUnique(input []string, search string) []string {
	for x := 0; x < len(input); x++ {
		if input[x] == search {
			return input
		}
	}
	return append(input, search)
}
func diagonalsSearch(input [][]string, search string) ([]string, int) {
	total := 0
	completeList := make([]string, 0)
	topRightList := make([]string, 0)
	bottomLeftList := make([]string, 0)
	topLeftList := make([]string, 0)
	bottomRightList := make([]string, 0)
	MASList := make([]string, 0)
	for z := 0; z < len(input); z++ {
		topRightAccumulator := ""
		bottomLeftAccumulator := ""
		topLeftAccumulator := ""
		bottomRightAccumulator := ""
		for x := 0; x < len(input); x++ {

			for y := 0; y < len(input); y++ {
				//Hardcoded MAS search
				if input[x][y] == "A" {
					coordinates := fmt.Sprintf("%d,%d", x, y)
					//Search around and X
					buffer := ""
					if x-1 >= 0 && y-1 >= 0 && (input[x-1][y-1] == "M" || input[x-1][y-1] == "S") { //topLeft
						buffer += input[x-1][y-1]
					}
					if x+1 < len(input) && y+1 < len(input) && (input[x+1][y+1] == "M" || input[x+1][y+1] == "S") { //bottomRight
						buffer += input[x+1][y+1]
					}
					if x+1 < len(input) && y-1 >= 0 && (input[x+1][y-1] == "M" || input[x+1][y-1] == "S") { //topRight
						buffer += input[x+1][y-1]
					}
					if x-1 >= 0 && y+1 < len(input) && (input[x-1][y+1] == "M" || input[x-1][y+1] == "S") { //bottomLeft
						buffer += input[x-1][y+1]
					}
					if len(buffer) == 4 && z == 0 {
						fmt.Println(buffer, coordinates)
						isMas := evaluateMAS(buffer)
						if isMas {
							MASList = appendIfUnique(MASList, coordinates)
						}
					}
				}

				if x+y+z == len(input)-1 && z != 0 {
					topLeftAccumulator += input[x][y]
					// fmt.Print("👍")
					// continue
				}
				if x+y-z == len(input)-1 {
					bottomRightAccumulator += input[x][y]
					// fmt.Print("🤑")
					// continue
				}

				if x == y+z && z != 0 {
					bottomLeftAccumulator += input[x][y]
					// fmt.Print("🤣")
					// continue
				}
				if y == x+z {
					topRightAccumulator += input[x][y]
					// fmt.Print("😊")
					// continue
				}
				// fmt.Print("⬛")
			}
			// fmt.Println()
		}
		topRightList = append(topRightList, topRightAccumulator)
		bottomLeftList = append(bottomLeftList, bottomLeftAccumulator)
		topLeftList = append(topLeftList, topLeftAccumulator)
		bottomRightList = append(bottomRightList, bottomRightAccumulator)
		// fmt.Println("////////")
	}
	// completeList = append(completeList, topRightList...)
	// completeList = append(completeList, bottomLeftList...)
	// completeList = append(completeList, topLeftList...)
	// completeList = append(completeList, bottomRightList...)
	fmt.Println(MASList, len(MASList), "MAS LIST")
	bottomLeftToRightTop := make([]string, 0)
	bottomLeftToRightTop = append(bottomLeftToRightTop, bottomLeftList...)
	bottomLeftToRightTop = append(bottomLeftToRightTop, topRightList...)

	topLeftToBottomRight := make([]string, 0)
	topLeftToBottomRight = append(topLeftToBottomRight, topLeftList...)
	topLeftToBottomRight = append(topLeftToBottomRight, bottomRightList...)


	completeList = append(completeList, bottomLeftToRightTop...)
	completeList = append(completeList, topLeftToBottomRight...)
	for _, item := range completeList {
		completeList = append(completeList, reverseWord(item))
	}
	total += countInArray(completeList, search)
	
	return completeList, total // remove the always 2 dupplicates
}
func main() {
	// dat, _ := os.ReadFile("./test4_3.txt")
	// dat, _ := os.ReadFile("./test3_2.txt")
	dat, _ := os.ReadFile("./input4_1.txt")
	searchWord := "MAS"
	stringValue := string(dat)
	matrixs := inputToMatrixs(stringValue)
	total := 0
	for i := 0; i < 4; i++ {
		val := inlineSearch(matrixs, searchWord) //plain
		matrixs = rotateArray(matrixs)
		total += val
	}
	fmt.Println(total)
	_, total2 := diagonalsSearch(matrixs, searchWord)
	fmt.Println(total + total2)
}
//> 1406