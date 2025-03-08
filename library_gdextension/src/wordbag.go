package main

import (
	"fmt"
	"math/rand/v2"

	"gdextension/wordlists"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Packed"
	"graphics.gd/variant/String"
)

type Wordbag struct {
	classdb.Extension[Wordbag, Node.Instance] `gd:"Wordbag"`
	CurIdx int64
	initialized bool
	wb wordlists.Wordbag
}

func RandomSeed() int64 {
	return rand.Int64()
}

func (w *Wordbag) Init(seed int64) {
	if w.initialized { return }

	Engine.Println(fmt.Sprintf("Created wordbag with seed %v!", seed))
	w.wb = wordlists.NewWordbag(2, &wordlists.EasyWords, &seed)

	w.initialized = true
}

func (w *Wordbag) GetWord(idx int64) String.Readable {
	if !w.initialized {
		Engine.Println("Tried to get a word from uninitialized wordbag!")
		return String.New()
	}
	if idx < 0 {
		Engine.Println("Tried to get wordbag word of index ", idx, "!")
		return String.New()
	}

	return String.New(w.wb.GetWord(int(idx)))
}

func (w *Wordbag) GetNext(count int64) Packed.Strings {
	if !w.initialized {
		Engine.Println("Tried to get a word from uninitialized wordbag!")
		return Packed.MakeStrings()
	}

	words := Packed.Strings{}

	for i := w.CurIdx; i < w.CurIdx + count; i++ {
		words.Append(String.New(w.GetWord(i)))
	}
	w.CurIdx += count

	return words
}

func (w *Wordbag) GetSeed() int64 {
	return w.wb.Seed
}