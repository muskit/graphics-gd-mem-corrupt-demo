package wordlists_test

import (
	"gdextension/wordlists"
	"testing"
)

func TestListUniqueness(t *testing.T) {
	freq := map[string]uint{}
	for _, w := range wordlists.EasyWords {
		freq[w] += 1
		if freq[w] > 1 {
			t.Errorf("Found extra occurrence of \"%v\"!", w)
		}
	}
}

func TestBagSimple(t *testing.T) {
	words := []string{"apple", "bang", "delta"}
	freq := map[string]uint{}
	bag := wordlists.NewWordbag(1, &words, nil)
	for i := 0; i < len(words); i++ {
		w := bag.GetWord(i)
		freq[w] += 1
	}

	// uniqueness check
	if len(freq) != len(words) {
		t.Errorf("Bag has %v unique words; should have %v!", len(freq), len(words))
	}

	// frequency check
	for _, f := range freq {
		if f != 1 {
			t.Errorf("Found duplicate word in bag!")
		}
	}
}

func TestBagMultipleOccurrence(t *testing.T) {
	occurrences := 3

	words := []string{"apple", "bang", "delta"}
	freq := map[string]uint{}
	bag := wordlists.NewWordbag(occurrences, &words, nil)
	for i := 0; i < len(words)*occurrences; i++ {
		w := bag.GetWord(i)
		freq[w] += 1
	}

	// uniqueness
	if len(freq) != len(words) {
		t.Errorf("Bag has %v unique words; should have %v!", len(freq), len(words))
	}

	// frequency
	for _, f := range freq {
		if f != uint(occurrences) {
			t.Errorf("Wrong # of occurrences in bag! (found %v, expected %v)", freq, occurrences)
		}
	}
}

func TestBagOOBIndexing(t *testing.T) {
	words := []string{"apple", "bang", "delta"}
	bag := wordlists.NewWordbag(1, &words, nil)
	freq := map[string]uint{}

	for i := 0; i < 2*len(words); i++ {
		w := bag.GetWord(i)
		freq[w] += 1
	}

	for _, f := range freq {
		if f != 2 {
			t.Errorf("Bag was not expanded properly!")
		}
	}
}