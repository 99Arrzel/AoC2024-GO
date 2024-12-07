package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func safeRange(x int, y int) (result bool) {
	maxGap := 3
	if x == y {
		return false
	}
	if (int(math.Abs(float64(x - y))) > maxGap) {
		return false
	}
	return true
}

func main() {
	//Read
	// dat, err := os.ReadFile("./test2_1.txt")
	dat, _ := os.ReadFile("./input2_1.txt")
	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	safeCounter := 0
	for _, rowValue := range rowValues {
		values := strings.Split(rowValue, " ")
		isSafe := true
		firstValue, _ := strconv.Atoi(values[0])
		secondValue, _ := strconv.Atoi(values[1])
		increasingOrDecreasing := "increasing"
		if firstValue == secondValue {
			continue
		}
		if firstValue > secondValue {
			increasingOrDecreasing = "decreasing"
		}

		for idx, value := range values {
			intValue, _ := strconv.Atoi(value)
			if idx > 0 {
				oldIntValue, _ := strconv.Atoi(values[idx-1])
				if !safeRange(oldIntValue, intValue) {
					fmt.Println("UNSAFE", intValue, oldIntValue)
					isSafe = false
				}
				if increasingOrDecreasing == "increasing" && oldIntValue > intValue {
					isSafe = false
					break
				}
				if increasingOrDecreasing == "decreasing" && intValue > oldIntValue {
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			safeCounter++
		}
	}
	fmt.Println(safeCounter)
}
