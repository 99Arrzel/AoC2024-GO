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
	fmt.Println(len(initialStones))
}

func blink(initialstones []int) []int {
	counters0 := 0
	countersmod := 0
	// countersOthers := 0
	result := make([]int, 0)
	for _, stone := range initialstones {
		if stone == 0 {
			counters0++
			// result = append(result, 1)
			continue
		}
		stoneLen := lenLoop10(stone)
		if stoneLen%2 == 0 {
			countersmod++
			res := strconv.Itoa(stone)
			leftPart, _ := strconv.Atoi(res[0:stoneLen/2])
			rightPart, _ := strconv.Atoi(res[stoneLen/2:])
			result = append(result, leftPart, rightPart)
			continue
		}
		result = append(result, stone*2024)
	}
	appendedOnes := make([]int, counters0)
	for idx := range appendedOnes {
		appendedOnes[idx] = 1
	}
	result = append(result, appendedOnes...)
	


	// fmt.Println("___________")
	// fmt.Println(counters0, countersmod, countersOthers, "Z")
	// fmt.Println("___________")
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
