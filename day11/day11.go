package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Test")
	initialStones := []int{6563348, 67, 395, 0, 6,4425,89567,739318}
	// initialStones := []int{125, 17}
	for i:= 0; i < 75; i++{
		fmt.Println(i, len(initialStones))
		blink(&initialStones)
		// fmt.Println(initialStones, "Iteration", i )
		
	}
	fmt.Println(len(initialStones))
}

func blink(initialstones *[]int) {
	initialLen := len(*initialstones)
	// for index := range initialstones {
	for i := 0; i < initialLen; i++ {
		didMutate:=processStone(initialstones, i)
		if didMutate {
			initialLen = len(*initialstones)
			i++
		}
	}
	
}

func processStone(initialStones *[]int, index int) bool {

	inIndexValue := (*initialStones)[index]
	if inIndexValue == 0 {
		(*initialStones)[index] = 1
		return false
	}
	res := strconv.Itoa((*initialStones)[index])
	if len(res)%2 == 0 {
		leftPart, _ := strconv.Atoi(res[0 : len(res)/2])
		rightPart, _ := strconv.Atoi(res[len(res)/2:])
		(*initialStones) = slices.Insert((*initialStones), index, leftPart)
		(*initialStones)[index+1] = rightPart

		return true
	}
	(*initialStones)[index] = (*initialStones)[index] * 2024
	return false
}
