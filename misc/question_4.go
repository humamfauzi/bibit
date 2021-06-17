package question

import (
	"fmt"
)

func Examples() {
	examples := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println("Hello, playground", CategorizeAnagram(examples))
}

func CategorizeAnagram(examples []string) [][]string {
	anagramMap := make(map[int][]string)
	for _, word := range examples {
		feedRunesToMap(anagramMap, word)
	}
	return composeResult(anagramMap)
}

func feedRunesToMap(anagramMap map[int][]string, word string) {
	mult := multiplyRune([]rune(word))
	anagramMap[mult] = append(anagramMap[mult], word)
}

func composeResult(anagramMap map[int][]string) [][]string {
	finalResult := make([][]string, len(anagramMap))
	index := 0
	for _, v := range anagramMap {
		finalResult[index] = v
		index++
	}
	return finalResult
}

func multiplyRune(runes []rune) int {
	result := 1
	for _, v := range runes {
		result *= int(v)
	}
	return result
}