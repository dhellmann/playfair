package cipher

import (
	"fmt"
	"strings"
)

type Matrix struct {
	keyword string
	content [][]string
}

const alphabet = "abcdefghiklmnopqrstuvwxyz"
const minKeywordLen = 6

func NewMatrix(keyword string) (*Matrix, error) {
	m := &Matrix{
		keyword: keyword,
	}
	m.content = make([][]string, 5)
	for i := 0; i < 5; i++ {
		m.content[i] = make([]string, 5)
	}

	uniq := ""
	for i := 0; i < len(keyword); i++ {
		c := keyword[i : i+1]
		if c == "j" { // treat "i" and "j" as the same letter
			c = "i"
		}
		if !strings.Contains(uniq, c) {
			uniq = uniq + keyword[i:i+1]
		}
	}
	if len(uniq) < minKeywordLen {
		return nil, fmt.Errorf("matrix keyword must have at least %d unique letters", minKeywordLen)
	}

	remaining := alphabet
	row := 0
	col := 0
	for i := 0; i < len(uniq); i++ {
		c := uniq[i : i+1]
		remaining = strings.Replace(remaining, c, "", -1)
		fmt.Printf("%d, %d: %s\n", row, col, c)
		m.content[row][col] = c
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	for i := 0; i < len(remaining); i++ {
		c := remaining[i : i+1]
		fmt.Printf("%d, %d: %s\n", row, col, c)
		m.content[row][col] = c
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	return m, nil
}
