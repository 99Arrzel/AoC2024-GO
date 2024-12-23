package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Test")
	initialStones := []int{6563348, 67, 395, 0, 6,4425,89567,739318}
	// initialStones := []int{125, 17}
	
	for i := 0; i< 75 ; i++{
		initialStones = blink(initialStones)
		fmt.Println(i, len(initialStones))
	}
	// fmt.Println(initialStones)	
	fmt.Println(len(initialStones))
}

func blink (initialstones []int) []int{
	result := make([]int,0)
	for _, stone := range initialstones {
		result = append(result, processStone(stone)...)
	}
	return result
}

func processStone(a int) []int {
	if a == 0 {
		return []int{1}
	}
	res := strconv.Itoa(a)
	if len(res)%2 == 0 {
		leftPart, _ := strconv.Atoi(res[0 : len(res)/2])
		rightPart, _ := strconv.Atoi(res[len(res)/2:])
		return []int{leftPart, rightPart}
	}
	return []int{a * 2024}
}
