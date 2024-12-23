package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Test")
	initialStones := []int{6563348, 67, 395, 0, 6, 4425, 89567, 739318}
	// initialStones := []int{125, 17}
	startTime := time.Now()
	for i := 0; i < 75; i++ {
		initialStones = blink(initialStones)
		timed := time.Since(startTime)
		fmt.Println(i, len(initialStones), "MS", timed.Milliseconds())
	}
	// fmt.Println(initialStones)
	fmt.Println(len(initialStones), initialStones)
}

func blink(initialstones []int) []int {
	counters0 := 0
	append2024Multiplies := make([]int, 0, len(initialstones))
	appendPairValues := make([]int, 0, len(initialstones))
	result := make([]int, 0)
	for idx, stone := range initialstones {
		if stone == 0 {
			counters0++
			continue
		}
		stoneLen := lenLoop10(stone)
		if stoneLen%2 == 0 {
			appendPairValues = append(appendPairValues, idx)
			// halfLen := stoneLen / 2
			// res := strconv.Itoa(stone)
			// leftPart, _ := strconv.Atoi(res[0:halfLen])
			// rightPart, _ := strconv.Atoi(res[halfLen:])
			// result = append(result, leftPart, rightPart)
			continue
		}
		// result = append(result, stone*2024)
		append2024Multiplies = append(append2024Multiplies, idx)
	}
	appendedOnes := make([]int, counters0)
	for idx := range appendedOnes {
		appendedOnes[idx] = 1
	}
	for idx := range append2024Multiplies {
		append2024Multiplies[idx] = initialstones[append2024Multiplies[idx]] * 2024
	}
	pairValues := make([]int, 0,len(appendPairValues)*2)
	for idx := range appendPairValues {
		stone := initialstones[appendPairValues[idx]]
		stoneLen := lenLoop10(stone)
		halfLen := stoneLen / 2
		res := strconv.Itoa(stone)
		leftPart, _ := strconv.Atoi(res[0:halfLen])
		rightPart, _ := strconv.Atoi(res[halfLen:])
		pairValues = append(pairValues, leftPart, rightPart)
	}

	result = append(result, appendedOnes...)
	result = append(result, append2024Multiplies...)
	result = append(result, pairValues...)
	return result
}

// https://stackoverflow.com/questions/68122675/fastest-way-to-find-number-of-digits-of-an-integer

func lenLoop10(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}
