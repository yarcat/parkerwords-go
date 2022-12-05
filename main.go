package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strings"
	"time"
)

type index struct {
	b2w   map[uint32]string
	index [26][]uint32
	abc   [26]int
}

func (i *index) word(w uint32) string { return i.b2w[w] }

func newIndexFromWords(data string) index {
	var freq [26]struct {
		c byte
		f int
	}
	for i := range freq {
		freq[i].c = byte(i)
	}
	b2w := make(map[uint32]string, 6000)
	for _, w := range strings.Fields(data) {
		if len(w) != 5 {
			continue
		}
		b := uint32(1<<(w[0]-'a') | 1<<(w[1]-'a') | 1<<(w[2]-'a') | 1<<(w[3]-'a') | 1<<(w[4]-'a'))
		if bits.OnesCount32(b) == 5 && b2w[b] == "" {
			b2w[b] = w
			for _, c := range w {
				freq[c-'a'].f++
			}
		}
	}
	sort.Slice(freq[:], func(i, j int) bool { return freq[i].f < freq[j].f })
	var abc, cba [26]int
	for i, f := range freq {
		abc[i] = int(f.c)
		cba[f.c] = i
	}
	var abcIndex [26][]uint32
	for b := range b2w {
		min := cba[bits.TrailingZeros32(b)]
		for m := b & (b - 1); m != 0; m &= m - 1 {
			if i := cba[bits.TrailingZeros32(m)]; i < min {
				min = i
			}
		}
		abcIndex[min] = append(abcIndex[min], b)
	}
	return index{b2w: b2w, index: abcIndex, abc: abc}
}

func find(index *index, i int, cur uint32, toSkip int, found []uint32, res [][5]uint32) [][5]uint32 {
	for i := i; toSkip >= 0 && i < 26; i++ {
		if cur&(1<<index.abc[i]) != 0 {
			continue
		}
		for _, w := range index.index[i] {
			if cur&w != 0 {
				continue
			}
			if f := append(found, w); len(f) == 5 {
				res = append(res, *(*[5]uint32)(f))
			} else {
				res = find(index, i+1, cur|w, toSkip, f, res)
			}
		}
		toSkip--
	}
	return res
}

func findAll(index *index) (res [][5]uint32) {
	var buf [5]uint32
	return find(index, 0 /*i*/, 0 /*cur*/, 1 /*toSkip*/, buf[:0], res)
}

func findAllPar(index *index) (res [][5]uint32) {
	jobs := len(index.index[0]) + len(index.index[1])
	out := make(chan [][5]uint32, jobs)
	for i := 0; i < 2; i++ {
		i := i
		for _, w := range index.index[i] {
			cur := w
			go func() {
				buf := [5]uint32{cur}
				out <- find(index, i+1, cur, 1-i /*toSkip*/, buf[:1], nil)
			}()
		}
	}
	for i := 0; i < jobs; i++ {
		res = append(res, <-out...)
	}
	return res
}

func main() {
	start := time.Now()
	wordsData, err := os.ReadFile("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	index := newIndexFromWords(string(wordsData))

	startAlgo := time.Now()
	res := findAllPar(&index)

	startWrite := time.Now()
	f, _ := os.Create("solutions.txt")
	out := bufio.NewWriter(f)
	for _, r := range res {
		for _, w := range r {
			out.WriteString(index.word(w))
			out.WriteByte('\t')
		}
		out.WriteByte('\n')
	}
	out.Flush()
	f.Close()

	end := time.Now()

	fmt.Printf("%d solutions written to solutions.txt\n", len(res))
	fmt.Printf("Total time:%6vµs\n", end.Sub(start).Microseconds())
	fmt.Printf("Read:      %6vµs\n", startAlgo.Sub(start).Microseconds())
	fmt.Printf("Process:   %6vµs\n", startWrite.Sub(startAlgo).Microseconds())
	fmt.Printf("Write:     %6vµs\n", end.Sub(startWrite).Microseconds())
}
