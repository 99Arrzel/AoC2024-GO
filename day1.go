package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Read
	dat, err := os.ReadFile("./input1_1.txt")
	check(err)

	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	firstColumn := make([]int32, 0)
	secondColumn := make([]int32, 0)
	for _, rowValue := range rowValues {
		values := strings.Split(rowValue, "   ")
		firstValue, err := strconv.Atoi(values[0])
		secondValue, err := strconv.Atoi(values[1])
		check(err)
		firstColumn = append(firstColumn, int32(firstValue))
		secondColumn = append(secondColumn, int32(secondValue))
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
}
