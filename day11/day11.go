package main

import (
	"fmt"
	"maps"
	"strconv"
)

func main() {
	// initialStones := []int{6563348, 67, 395, 0, 6, 4425, 89567, 739318}
	initialStones := []int{125, 17}
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
	blink(initialStones)
	for i := 0; i < 25; i++ {
		// timed := time.Since(startTime)
		// fmt.Println(i, len(initialStones), "MS", timed.Milliseconds())
	}
	// fmt.Println(initialStones)
	fmt.Println(len(initialStones), initialStones)
}

func blink(initialstones []int) (map[int]int, int) {
	mappedCounters := make(map[int]int)
	for _, stone := range initialstones {
		_, ok := mappedCounters[stone]
		// mappedCounters[stone]++
		if ok {
			mappedCounters[stone]++
			continue
		}
		mappedCounters[stone] = 1
	}


	counter := 0
	mappedKeys := maps.Keys(mappedCounters)
	fmt.Println("Before", mappedCounters)
	mappedResult := make(map[int]int)
	for key := range mappedKeys {
		// value := mappedCounters[key]
		newKeys := processStone(key)
		for _, keyToAdd := range newKeys {
			val, ok := mappedResult[keyToAdd]
			if ok {
				counter += val
				mappedResult[keyToAdd]++
				continue
			}
			mappedResult[keyToAdd] = 1
			counter++
		}
	}
	fmt.Println(mappedResult, counter)
	return mappedResult, counter
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
	res := strconv.Itoa(stone)
	if len(res)%2 == 0 {
		leftPart, _ := strconv.Atoi(res[0 : len(res)/2])
		rightPart, _ := strconv.Atoi(res[len(res)/2:])
		return []int{leftPart, rightPart}
	}
	return []int{stone * 2024}
}
