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
	if int(math.Abs(float64(x-y))) > maxGap {
		return false
	}
	return true
}
func removeAndNew(slice []int, s int) []int {
	newSlice := make([]int, 0)
	for idx := 0; idx < len(slice); idx++ {
		if idx == s {
			continue
		}
		newSlice = append(newSlice, slice[idx])
	}
	return newSlice
}

// Must be of at least lenght 2
func evaluateRowSafeness(values []int) bool {
	arrLen := len(values)
	if arrLen < 2 {
		return true
	}
	firstValue := values[0]
	secondValue := values[1]
	increasingOrDecreasing := "increasing"
	if firstValue > secondValue {
		increasingOrDecreasing = "decreasing"
	}
	for idx := 0; idx < arrLen; idx++ {
		intValue := values[idx]
		if idx > 0 {
			oldIntValue := values[idx-1]
			if !safeRange(oldIntValue, intValue) {
				fmt.Println("Not safe by range", oldIntValue, intValue)
				return false
			} else if increasingOrDecreasing == "increasing" && oldIntValue > intValue {
				fmt.Println("Not safe due to change of increase", oldIntValue, intValue)
				return false
			} else if increasingOrDecreasing == "decreasing" && intValue > oldIntValue {
				fmt.Println("Not safe due to change of decrease", oldIntValue, intValue)
				return false
			}
		}
	}
	return true
}
func toIntSlice(values []string) []int {
	newSlice := make([]int, 0)
	for idx := 0; idx < len(values); idx++ {
		intVal, _ := strconv.Atoi(values[idx])
		newSlice = append(newSlice, intVal)
	}
	return newSlice
}

func main() {
	//Read
	// dat, _ := os.ReadFile("./test2_1.txt")
	dat, _ := os.ReadFile("./input2_1.txt")
	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	safeCounter := 0

	for _, rowValue := range rowValues {
		values := toIntSlice(strings.Split(rowValue, " "))
		isSafe := evaluateRowSafeness(values)
		if isSafe {
			safeCounter++
		} else {
			//Try removing index positions
			for idx := 0; idx < len(values); idx++ {
				newSlice := removeAndNew(values, idx)
				fmt.Println(newSlice)
				isReallySafe := evaluateRowSafeness(newSlice)
				if isReallySafe {
					safeCounter++
					break
				}
			}
		}

	}
	fmt.Println(safeCounter, "Number of safe rows")
}
