package main

import (
	"fmt"
	"maps"
	"strconv"
)

func main() {
	initialStones := []int{6563348, 67, 395, 0, 6, 4425, 89567, 739318}
	// initialStones := []int{125, 17}
	mappedCounters := make(map[int]int)
	for _, stone := range initialStones {
		_, ok := mappedCounters[stone]
		// mappedCounters[stone]++
		if ok {
			mappedCounters[stone]++
			continue
		}
		mappedCounters[stone] = 1
	}
	// startTime := time.Now()
	counter := 0
	for i := 0; i < 75; i++ {
		stones := blink(mappedCounters)
		mappedCounters = stones
	}
	for key := range maps.Keys(mappedCounters) {
		counter += mappedCounters[key]
	}
	fmt.Println("Final", counter)
}

func blink(mappedCounters map[int]int) map[int]int {
	counter := 0
	mappedKeys := maps.Keys(mappedCounters)
	mappedResult := make(map[int]int)
	for key := range mappedKeys {
		newKeys := processStone(key)
		val1 := mappedCounters[key]
			for _, keyToAdd := range newKeys {
				_, ok := mappedResult[keyToAdd]
				if !ok {
					mappedResult[keyToAdd] = val1
					counter++
					continue
				}
				mappedResult[keyToAdd] += val1
				counter += val1
			}
	}

	return mappedResult
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

func processStone(stone int) []int {

	if stone == 0 {
		return []int{1}
	}
	if stone != 1 && lenLoop10(stone)%2 == 0 {
		res := strconv.Itoa(stone)
		leftPart, _ := strconv.Atoi(res[0 : len(res)/2])
		rightPart, _ := strconv.Atoi(res[len(res)/2:])
		return []int{leftPart, rightPart}
	}
	return []int{stone * 2024}
}
