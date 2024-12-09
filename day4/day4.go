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
func diagonalsSearch(input [][]string, search string) ([]string, int) {
	total := 0
	completeList := make([]string, 0)
	topRightList := make([]string, 0)
	bottomLeftList := make([]string, 0)
	topLeftList := make([]string, 0)
	bottomRightList := make([]string, 0)
	for z := 0; z < len(input); z++ {
		topRightAccumulator := ""
		bottomLeftAccumulator := ""
		topLeftAccumulator := ""
		bottomRightAccumulator := ""
		for x := 0; x < len(input); x++ {
			for y := 0; y < len(input); y++ {
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
	completeList = append(completeList, topRightList...)
	completeList = append(completeList, bottomLeftList...)
	completeList = append(completeList, topLeftList...)
	completeList = append(completeList, bottomRightList...)
	for _, item := range completeList {
		completeList = append(completeList, reverseWord(item))
	}
	total += countInArray(completeList, search)
	fmt.Println("Total", total)
	fmt.Println(completeList)
	return completeList, total // remove the always 2 dupplicates
}
func main() {
	// dat, _ := os.ReadFile("./test4_1.txt")
	// dat, _ := os.ReadFile("./test3_2.txt")
	dat, _ := os.ReadFile("./input4_1.txt")
	searchWord := "XMAS"
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
