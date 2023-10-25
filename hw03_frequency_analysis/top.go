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

	pairs := rankByWordCountSliceAsterisk(str2)

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

func rankByWordCountSliceAsterisk(str string) []Pair {
	r := regexp.MustCompile(`[А-Яа-я]+(\-?[А-Яа-я]+)*`)
	words := r.FindAllString(str, -1)

	wordsKey := make(map[string]int)
	for _, word := range words {
		wordsKey[strings.ToLower(word)]++
	}
	delete(wordsKey, "-")

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
