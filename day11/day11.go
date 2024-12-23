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
		fmt.Println(i, len(initialStones), "MS", timed.Milliseconds() )
	}
	// fmt.Println(initialStones)
	fmt.Println(len(initialStones))
}

func blink(initialstones []int) []int {
	result := make([]int, 0)
	for _, stone := range initialstones {
		result = append(result, processStone(stone)...)
	}
	return result
}

func processStone(a int) []int {
	if a == 0 {
		return []int{1}
	}
	if lenLoop10(a)%2 == 0 {
		res := strconv.Itoa(a)
		leftPart, _ := strconv.Atoi(res[0 : len(res)/2])
		rightPart, _ := strconv.Atoi(res[len(res)/2:])
		return []int{leftPart, rightPart}
	}
	return []int{a * 2024}
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
