package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	fileName := flag.String("f", "word.list", "File Name")
	flag.Parse()

	// Read input
	words, err := readLines(*fileName)
	if err != nil {
		fmt.Errorf("File %s list does not exist.", fileName)
	}

	fmt.Printf("File Length: %d\n", len(words))

	// Sort string since the search is based on sorting
	sort.Strings(words)

	// Find the longest compound word
	start := time.Now()
	longestCompoundWord := FindLongestCompoundWord(words)
	fmt.Printf("Time %v Word: %s\n", time.Since(start), longestCompoundWord)
}

func FindLongestCompoundWord(words []string) string {
	longestCompoundWord := ""
	for _, word := range words {
		if len(word) < len(longestCompoundWord) {
			continue
		}

		// If the word is made out of 2 different words make that our current
		// longest compound word
		if isCompound(words, word) {
			longestCompoundWord = word
		}
	}
	return longestCompoundWord
}

func isCompound(words []string, word string) bool {
	for i, _ := range word {
		prefix := word[:i]
		postfix := word[i:]

		if len(prefix) == 0 || len(postfix) == 0 {
			continue
		}

		if !contains(words, prefix) {
			continue
		}

		// Recursively check if postfix is also a compound word
		if contains(words, postfix) || isCompound(words, postfix) {
			return true
		}
	}
	return false
}

func contains(a []string, x string) bool {
	return a[sort.SearchStrings(a, x)] == x
}

func readLines(fileName string) ([]string, error) {
	var lines []string

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
