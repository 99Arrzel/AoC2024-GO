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
		if strings.Contains(strings.Join(input[x], ""), search) {
			total++
		}
	}
	return total
}

func main() {
	//Read
	dat, _ := os.ReadFile("./test4_1.txt")
	// dat, _ := os.ReadFile("./test3_2.txt")
	// dat, _ := os.ReadFile("./input4_1.txt")
	searchWord := "XMAS"

	stringValue := string(dat)
	matrixs := inputToMatrixs(stringValue)
	flat := inlineSearch(matrixs, searchWord) //plain
	fmt.Println(flat, matrixs)
	rotated := rotateArray(matrixs)
	degrees90 := inlineSearch(rotated, searchWord) //plain
	fmt.Println(degrees90, rotated)
	rotated2 := rotateArray(rotated)
	degrees180 := inlineSearch(rotated2, searchWord) //plain
	fmt.Println(degrees180, rotated2)
	rotated3 := rotateArray(rotated2)
	degrees270 := inlineSearch(rotated3, searchWord) //plain
	fmt.Println(degrees270, rotated3)

	fmt.Println(flat + degrees90 + degrees180 + degrees270)
	// searchWord := "XMAS"

}
