package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//Read
	// dat, err := os.ReadFile("./test1_1.txt")
	dat, _ := os.ReadFile("./input1_1.txt")
	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	firstColumn := make([]int, 0)
	secondColumn := make([]int, 0)
	firstColumnMap := make(map[int]int)
	secondColumnMap := make(map[int]int)
	for _, rowValue := range rowValues {
		values := strings.Split(rowValue, "   ")
		firstValue, _ := strconv.Atoi(values[0])
		secondValue, _ := strconv.Atoi(values[1])
		_, firstMapOk := firstColumnMap[firstValue]
		_, secondMapOk := secondColumnMap[secondValue]
		if firstMapOk {
			firstColumnMap[firstValue]++
		} else {
			firstColumnMap[firstValue] = 1
		}
		if secondMapOk {
			secondColumnMap[secondValue]++
		} else {
			secondColumnMap[secondValue] = 1
		}
		firstColumn = append(firstColumn, firstValue)
		secondColumn = append(secondColumn, secondValue)
	}
	//Sort
	slices.Sort(firstColumn)
	slices.Sort(secondColumn)
	//Both arrays always same size
	totalDiff := 0
	for i := 0; i < len(firstColumn); i++ {
		totalDiff = totalDiff + int(math.Abs(float64(firstColumn[i])-float64(secondColumn[i])))
	}
	fmt.Println(totalDiff)
	//part 2
	similarity := 0
	for i := 0; i < len(firstColumn); i++ {
		multiplier := 0
		valFound, found := secondColumnMap[firstColumn[i]]
		if found {
			multiplier = valFound
		}
		similarity = similarity + multiplier*firstColumn[i]
	}
	fmt.Println(similarity)
}
