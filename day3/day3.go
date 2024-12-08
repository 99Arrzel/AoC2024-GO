package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read
	// dat, _ := os.ReadFile("./test3_1.txt")
	dat, _ := os.ReadFile("./input3_1.txt")
	stringValue := string(dat)
	rowValues := strings.Split(stringValue, "\n")
	// result := 0
	muls := make([][]int, 0)
	for _, rowValue := range rowValues {
		mulValues := strings.Split(rowValue, "mul(")
		for _, posibleMulValue := range mulValues {
			toEvaluateMul := strings.Split(posibleMulValue, ",")
			if len(toEvaluateMul) < 2 {
				fmt.Println("Nothing to evalute")
				continue
			}
			//validating first number
			leftValueNumber, err := strconv.Atoi(toEvaluateMul[0])
			if err != nil {
				fmt.Println("Left number of mul not valid", toEvaluateMul[0])
				continue
			}
			//Second
			if len(toEvaluateMul) > 1 {
				rightSideSplit := strings.Split(toEvaluateMul[1], ")")
				//We only care about first value
				rightValueNumber, err2 := strconv.Atoi(rightSideSplit[0])
				fmt.Println(rightValueNumber)
				if err2 == nil {
					values := make([]int, 2)
					values[0] = leftValueNumber
					values[1] = rightValueNumber
					muls = append(muls, values)
				}
				//Only results with
			}
		}
	}
	//Add multipliyers
	total := 0
	for _, rowValue := range muls {
		total = total + (rowValue[0] * rowValue[1])
	}
	fmt.Println(muls)
	fmt.Println(total)
}
