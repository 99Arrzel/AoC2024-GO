package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// rules, _ := os.ReadFile("./test_rules.txt")
	rules, _ := os.ReadFile("./input_rules.txt")
	linesRules := strings.Split(string(rules), "\n")
	// manual, _ := os.ReadFile("./test_manual.txt")
	manual, _ := os.ReadFile("./input_manual.txt")
	linesManual := strings.Split(string(manual), "\n")
	validLinesManual := make([]string, 0)
	invalidLinesManual := make([]string, 0)
	for _, line := range linesManual {
		matchedRules := matchRulesToLine(linesRules, line)
		isValidLine := isValid(matchedRules, line)
		if isValidLine {
			validLinesManual = append(validLinesManual, line)
		} else {
			invalidLinesManual = append(invalidLinesManual, line)
		}
	}
	//Re order invalid
	reEvaluatedInvalidLines := make([]string, 0)
	for _, line := range invalidLinesManual {
		matchedRules := matchRulesToLine(linesRules, line)
		newValidLine := string(line)
		isValidLine := isValid(matchedRules, line)
		swapCounter := 10
		indexes := findIndexOfRules(matchedRules, line)
		for !isValidLine && swapCounter > 0 {
			// fmt.Println(indexes, line)
			for x := 0; x < len(indexes); x = x + 2 {
				if indexes[x] > indexes[x+1] {
					newValidLine = swapRules(newValidLine, indexes[x], indexes[x+1])
					isValidLine = isValid(matchedRules, newValidLine)
					indexes = findIndexOfRules(matchedRules, newValidLine)
					if isValidLine {
						fmt.Println("Is valid", newValidLine)
						break
					}
				}
			}
			swapCounter--
		}
		if swapCounter == 0 {
			fmt.Println("No solution found")
		}
		reEvaluatedInvalidLines = append(reEvaluatedInvalidLines, newValidLine)
	}

	totalValid := sumMiddleValuesInManualLines(validLinesManual)
	totalInvalid := sumMiddleValuesInManualLines(invalidLinesManual)
	totalReevaluated := sumMiddleValuesInManualLines(reEvaluatedInvalidLines)
	fmt.Println(reEvaluatedInvalidLines)
	fmt.Println(totalValid, totalInvalid, totalReevaluated)
}

func sumMiddleValuesInManualLines(invalidLinesManual []string) int {
	total := 0
	for _, line := range invalidLinesManual {
		items := strings.Split(line, ",")
		item := items[len(items)/2]
		val, _ := strconv.Atoi(item)
		total += val
	}
	return total
}

func matchRulesToLine(linesRules []string, line string) []string {
	matchedRules := make([]string, 0)
	for _, lineRule := range linesRules {
		rules := strings.Split(lineRule, "|")
		matchesThisLineRule := true

		if !strings.Contains(line, rules[0]) || !strings.Contains(line, rules[1]) {
			matchesThisLineRule = false
		}

		if matchesThisLineRule {
			matchedRules = append(matchedRules, lineRule)
		}
	}
	return matchedRules
}

func isValid(matchedRules []string, line string) bool {
	indexes := findIndexOfRules(matchedRules, line)
	for x := 0; x < len(indexes); x = x + 2 {
		if indexes[x] > indexes[x+1] {
			return false
		}
	}
	return true
}


func findIndexOfRules(matchedRules []string, line string) []int {
	indexes := make([]int, 0)
	for _, rule := range matchedRules {
		rules := strings.Split(rule, "|")
		indexOfFirstRule := strings.Index(line, rules[0])
		indexOfSecondRule := strings.Index(line, rules[1])
		if indexOfFirstRule > indexOfSecondRule {
			indexes = append(indexes, indexOfFirstRule)
			indexes = append(indexes, indexOfSecondRule)
		}
	}
	return indexes
}
func swapRules(line string, index1 int, index2 int) string {
	//rule always 2 characters
	newStr := ""
	for x := 0; x < len(line); x++ {
		if x == index1 {
			newStr += string(line[index2])
			continue
		}
		if x == index2 {
			newStr += string(line[index1])
			continue
		}
		if x == index1+1 {
			newStr += string(line[index2+1])
			continue
		}
		if x == index2+1 {
			newStr += string(line[index1+1])
			continue
		}
		newStr += string(line[x])
	}
	return newStr
}
