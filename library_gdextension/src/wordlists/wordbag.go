package wordlists

import "math/rand"

type Wordbag struct {
	Seed int64
	// how many times a word may occur before having to wait for the next bag "cycle"
	occurrences int
	idxCache []int
	rng rand.Rand
	wordlist *[]string
}

func NewWordbag(occurrences int, words *[]string, seed *int64) Wordbag {
	if seed == nil {
		s := int64(rand.Uint64())
		seed = &s
	}
	w := Wordbag{
		Seed: int64(rand.Uint64()),
		occurrences: max(1, occurrences),
		rng: *rand.New(rand.NewSource(*seed)),
		idxCache: make([]int, 0),
		wordlist: words,
	}
	
	w.extendBag()
	return w
}

func (w *Wordbag) extendBag() {
	listLen := len(*w.wordlist)
	size := w.occurrences * listLen

	ext := make([]int, size)
	for i := 0; i < w.occurrences; i++ {
		for j := 0; j < len(*w.wordlist); j++ {
			ext[i*listLen + j] = j
		}
	}

	w.rng.Shuffle(len(ext), func(i int, j int) {
		ext[i], ext[j] = ext[j], ext[i]
	})
	w.idxCache = append(w.idxCache, ext...)
}

func (w *Wordbag) GetWord(i int) string {
	// extend cache if idx falls outside capacity
	for i >= len(w.idxCache) {
		w.extendBag()
	}
	return (*w.wordlist)[w.idxCache[i]]
}

