package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkNoTrimAtoi(str string, intval int) bool {
	strval := strconv.Itoa(intval)
	return str != strval
}

func evaluateMuls(slice string) (result [][]int, total int) {
	posibleMuls := make([][]int, 0)
	tot := 0
	mulValues := strings.Split(slice, "mul(")
	for _, posibleMulValue := range mulValues {
		toEvaluateMul := strings.Split(posibleMulValue, ",")
		if len(toEvaluateMul) < 2 {
			// fmt.Println("Nothing to evalute", toEvaluateMul)
			continue
		}
		//validating first number
		leftValueNumber, err := strconv.Atoi(toEvaluateMul[0])
		if checkNoTrimAtoi(toEvaluateMul[0], leftValueNumber) {
			// fmt.Println("Might have spaces", toEvaluateMul)
			continue
		}
		if err != nil {
			// fmt.Println("Error evaluating left number", toEvaluateMul)
			continue
		}
		//Second
		if len(toEvaluateMul) > 1 {
			rightSideSplit := strings.Split(toEvaluateMul[1], ")")
			//We only care about first value
			rightValueNumber, err2 := strconv.Atoi(rightSideSplit[0])
			if checkNoTrimAtoi(rightSideSplit[0], rightValueNumber) {
				continue
			}

			if err2 != nil {
				// fmt.Println("Error evaluating right number", rightSideSplit[0])
				continue
			}
			//Check that Atoi is not trimming it

			if leftValueNumber != 0 && rightValueNumber != 0 {
				mul := make([]int, 2)
				mul[0] = leftValueNumber
				mul[1] = rightValueNumber
				tot += leftValueNumber * rightValueNumber
				posibleMuls = append(posibleMuls, mul)
			}
		}
	}
	return posibleMuls, tot
}

func main() {
	//Read
	// dat, _ := os.ReadFile("./test3_1.txt")
	// dat, _ := os.ReadFile("./test3_2.txt")
	dat, _ := os.ReadFile("./input3_1.txt")
	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	// result := 0
	muls := make([][][]int, 0)
	noSearchInstruction := "don't()"
	searchInstruction := "do()"
	// collector := ""
	// shouldAppend := true
	total := 0
	for idx, rowValue := range rowValues {
		letters := strings.Split(rowValue, "")
		collector := ""
		continuous := ""
		shouldAppend := true
		for _, letter := range letters {
			continuous += letter
			if strings.HasSuffix(continuous, noSearchInstruction) {
				shouldAppend = false
			}
			if strings.HasSuffix(continuous, searchInstruction) {
				shouldAppend = true
			}
			if shouldAppend {
				collector += letter
			}
		}
		rowMuls, tot := evaluateMuls(collector)
		total += tot
		fmt.Println(collector, idx, " ----------- ", tot)
		muls = append(muls, rowMuls)
	}
	fmt.Println(total)
}
