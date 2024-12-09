package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, _ := os.ReadFile("./input_rules.txt")
	linesRules := strings.Split(string(rules), "\n")
	manual, _ := os.ReadFile("./input_manual.txt")
	linesManual := strings.Split(string(manual), "\n")
	validLinesManual := make([]string, 0)
	for _, line := range linesManual {
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
		isValidLine := true
		for _, rule := range matchedRules {
			rules := strings.Split(rule, "|")
			indexOfFirstRule := strings.Index(line, rules[0])
			indexOfSecondRule := strings.Index(line, rules[1])
			if indexOfFirstRule > indexOfSecondRule {
				isValidLine = false
			}
		}
		if isValidLine {
			validLinesManual = append(validLinesManual, line)
		}
	}

	fmt.Println(len(validLinesManual), validLinesManual)
	
	total := 0
	for _, line := range validLinesManual {
		items := strings.Split(line, ",")
				item := items[len(items)/2] //Doubt if it works
		val, _ := strconv.Atoi(item)

		total += val
	}
	fmt.Println(total)
}
