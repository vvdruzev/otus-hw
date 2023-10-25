package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	if str == "" {
		return []string{}
	}

	r := regexp.MustCompile(`\s+`)
	str2 := r.ReplaceAllString(str, " ")

	pairs := rankByWordCountSlice(str2)

	keys := []string{}
	for _, pair := range pairs[:10] {
		keys = append(keys, pair.Key)
	}
	return keys
}

type Pair struct {
	Key   string
	Value int
}

func rankByWordCountSlice(str string) []Pair {
	wordsKey := make(map[string]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		wordsKey[word]++
	}

	ss := []Pair{}
	for k, v := range wordsKey {
		ss = append(ss, Pair{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value < ss[j].Value {
			return false
		}
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return true
	})

	return ss
}
