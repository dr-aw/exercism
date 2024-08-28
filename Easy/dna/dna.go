package dna

import (
	"fmt"
	"strings"
)

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (h Histogram, err error) {
	d = DNA(strings.ToUpper(string(d)))
	h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nuc := range d {
		switch nuc {
		case 'A', 'C', 'G', 'T':
			h[nuc]++
		default:
			return nil, fmt.Errorf("invalid nucleotide %v", nuc)
		}
	}
	return h, nil
}
