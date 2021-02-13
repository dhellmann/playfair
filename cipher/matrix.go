package cipher

import (
	"fmt"
	"strings"
)

// Matrix is an encryption/decryption tool
type Matrix struct {
	keyword string
	content [5][5]rune
}

func (m *Matrix) String() string {
	s := ""
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			s = s + string(m.content[row][col]) + " "
		}
		s = s + "\n"
	}
	return s
}

const alphabet = "abcdefghiklmnopqrstuvwxyz"
const minKeywordLen = 6

// NewMatrix creates a Matrix using the given keyword, or returns an error if the word cannot be used to create a matrix.
func NewMatrix(keyword string) (*Matrix, error) {
	keyword = strings.ToLower(keyword)
	m := &Matrix{
		keyword: keyword,
	}

	uniq := ""
	for _, r := range keyword {
		if r == 'j' { // treat "i" and "j" as the same letter
			r = 'i'
		}
		if !strings.Contains(alphabet, string(r)) {
			return nil, fmt.Errorf("matrix keywords must contain lower case ASCII letters only, '%c' is not valid", r)
		}
		if !strings.Contains(uniq, string(r)) {
			uniq = uniq + string(r)
		}
	}
	if len(uniq) < minKeywordLen {
		return nil, fmt.Errorf("matrix keyword must have at least %d unique letters", minKeywordLen)
	}

	remaining := alphabet
	row := 0
	col := 0
	for _, r := range uniq {
		c := string(r)
		remaining = strings.Replace(remaining, c, "", -1)
		m.content[row][col] = r
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	for _, r := range remaining {
		m.content[row][col] = r
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	return m, nil
}
