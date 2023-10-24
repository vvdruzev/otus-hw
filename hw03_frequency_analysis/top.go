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

	// pairs := rankByWordCount(str2)
	// pairs := rankByWordCountSlice(str2)
	pairs := rankByWordCountSliceAsterisk(str2)

	keys := []string{}
	for _, pair := range pairs[:10] {
		keys = append(keys, pair.Key)
	}
	return keys
}

func rankByWordCount(str string) PairList {
	wordsKey := make(map[string]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		if _, ok := wordsKey[word]; !ok {
			wordsKey[word] = 0
		}
		wordsKey[word]++
	}

	pl := make(PairList, len(wordsKey))
	i := 0
	for k, v := range wordsKey {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value > p[j].Value {
		return false
	}
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return true
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func rankByWordCountSlice(str string) []Pair {
	wordsKey := make(map[string]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		if _, ok := wordsKey[word]; !ok {
			wordsKey[word] = 0
		}
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

func rankByWordCountSliceAsterisk(str string) []Pair {
	r := regexp.MustCompile(`[\!,\,,\.,\?,\:]+`)
	str2 := r.ReplaceAllString(str, "")

	wordsKey := make(map[string]int)
	words := strings.Split(str2, " ")
	for _, word := range words {
		if _, ok := wordsKey[strings.ToLower(word)]; !ok {
			wordsKey[strings.ToLower(word)] = 0
		}
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
