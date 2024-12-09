package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseArray(slice []string) []string {
	newSlice := make([]string, 0)
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}
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
		newMatrixs[x] = reverseArray(newMatrixs[x])
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
			for y := 1; y < len(input); y++ {
				if x+y+z == len(input) {
					topLeftAccumulator += input[x][y]
				}
				if x+y-z == len(input) {
					bottomRightAccumulator += input[x][y]
				}
				if x == y+z {
					bottomLeftAccumulator += input[x][y]
				}
				if y == x+z {
					topRightAccumulator += input[x][y]
				}
			}
		}
		topRightList = append(topRightList, topRightAccumulator)
		bottomLeftList = append(bottomLeftList, bottomLeftAccumulator)
		topLeftList = append(topLeftList, topLeftAccumulator)
		bottomRightList = append(bottomRightList, bottomRightAccumulator)
	}
	completeList = append(completeList, topRightList...)
	completeList = append(completeList, bottomLeftList...)
	completeList = append(completeList, topLeftList...)
	completeList = append(completeList, bottomRightList...)
	for _, item := range completeList {
		reversed := reverseWord(item)
		completeList = append(completeList, reversed)
	}
	total += countInArray(completeList, search)
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
	list, total2 := diagonalsSearch(matrixs, searchWord)
	fmt.Println(list)
	fmt.Println(total + total2)
}
