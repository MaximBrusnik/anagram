package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)
	signatureToFirstWord := make(map[string]string)
	uniqueWords := make(map[string]map[string]struct{})

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortString(lowerWord)

		firstWord, exists := signatureToFirstWord[sortedWord]
		if !exists {
			signatureToFirstWord[sortedWord] = lowerWord
			firstWord = lowerWord
			uniqueWords[firstWord] = make(map[string]struct{})
		}

		// Проверяем, было ли слово уже добавлено
		if _, found := uniqueWords[firstWord][lowerWord]; !found {
			uniqueWords[firstWord][lowerWord] = struct{}{}
			anagramSets[firstWord] = append(anagramSets[firstWord], lowerWord)
		}
	}

	// Удаляем одиночные множества и сортируем
	result := make(map[string][]string)
	for firstWord, wordsInGroup := range anagramSets {
		if len(wordsInGroup) > 1 {
			sort.Strings(wordsInGroup)
			result[firstWord] = wordsInGroup
		}
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	// Пример входных данных
	inputWords := []string{"пятак", "Пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	// Находим анаграммы
	anagrams := findAnagrams(inputWords)

	// Выводим результат
	for key, group := range anagrams {
		fmt.Printf("%q: %v\n", key, group)
	}
}
