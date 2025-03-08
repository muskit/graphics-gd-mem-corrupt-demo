package wordlists

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed english_1k.txt
var words1k string
var EasyWords []string

//go:embed english_10k.txt
var words10k string
var MediumWords []string

// TODO
var painfulWords []string

func init() {
	populate(&EasyWords, words1k)
	populate(&MediumWords, words10k)
}

func populate(list *[]string, wordLines string) {
	lines := strings.Split(wordLines, "\n")
	for _, line := range lines {
		if len(line) != 0 {
			*list = append(*list, line)
		}
	}
}

func GetEasyWord() string {
	return EasyWords[rand.Int() % len(EasyWords)]
}

func GetMediumWord() string {
	return MediumWords[rand.Int() % len(MediumWords)]
}