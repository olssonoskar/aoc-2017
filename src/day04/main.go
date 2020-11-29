package main

import (
	"sort"
	"strings"
	"util"
)

func main() {
	input := util.GetInput("day04")
	println(part1(input))
	println(part2(input))
}

func part1(input string) int {
	validPhrase := 0
	phrases := strings.Split(input, "\n")
	for _, phrase := range phrases {
		if containsDuplicates(phrase) {
			continue
		}
		validPhrase++
	}
	return validPhrase
}

func part2(input string) int {
	validPhrase := 0
	phrases := strings.Split(input, "\n")
	for _, phrase := range phrases {
		if containsDuplicates(phrase) || containsAnagram(phrase) {
			continue
		}
		validPhrase++
	}
	return validPhrase
}

func containsDuplicates(phrase string) bool {
	wordCount := make(map[string]int)
	for _, word := range strings.Fields(phrase) {
		wordCount[word]++
		if wordCount[word] > 1 {
			return true
		}
	}
	return false
}

func containsAnagram(phrase string) bool {
	sortedCount := make(map[string]int)
	for _, word := range strings.Fields(phrase) {
		result := sorted(word)
		sortedCount[result]++
		if sortedCount[result] > 1 {
			return true
		}
	}
	return false
}

func sorted(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
