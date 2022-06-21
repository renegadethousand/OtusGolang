package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type word struct {
	name  string
	count int
}

func Top10(str string) []string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, ".", "")
	strSlice := strings.Fields(str)
	wordSlice := []word{}
	returnSlice := []string{}
	mapWords := map[string]int{}
	for _, v := range strSlice {
		if v == "-" {
			continue
		}
		mapWords[v]++
	}
	for key, value := range mapWords {
		wordSlice = append(wordSlice, word{name: key, count: value})
	}
	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].count != wordSlice[j].count {
			return wordSlice[i].count > wordSlice[j].count
		}
		return wordSlice[i].name < wordSlice[j].name
	})
	maxWords := 10
	for i, v := range wordSlice {
		if i >= maxWords {
			break
		}
		returnSlice = append(returnSlice, v.name)
	}
	return returnSlice
}
