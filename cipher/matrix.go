package cipher

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const alphabet = "abcdefghiklmnopqrstuvwxyz"
const minKeywordLen = 6

type location struct {
	row int
	col int
}

// Cipher is an encryption/decryption tool
type Cipher struct {
	keyword   string
	content   [5][5]rune
	locations map[rune]location
}

// Matrix returns the string representation of the cipher matrix
func (m *Cipher) Matrix() string {
	s := ""
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			s = s + string(m.content[row][col]) + " "
		}
		s = s + "\n"
	}
	return s
}

// nextValidRune starts at the beginning of the string and finds 1
// rune that is in the alphabet. If the rune is 'j' it is replaced
// with 'i'. If the input contains no valid runes, a space is
// returned.
func nextValidRune(input string) (rune, int) {
	var r rune
	var size int
	consumed := 0
	for {
		r, size = utf8.DecodeRuneInString(input[consumed:])
		if r == 'j' {
			r = 'i'
		}
		if size == 0 {
			return ' ', consumed
		}
		consumed += size
		if strings.Contains(alphabet, string(r)) {
			break
		}
	}
	return r, consumed
}

// runePairs returns a slice of slices of 2 runes for encodable
// characters from the alphabet
//
// If duplicate values are encountered, an 'x' is inserted between them.
func runePairs(input string) [][]rune {
	input = strings.ToLower(input)
	result := [][]rune{}
	consumed := 0
	for {
		var a, b rune
		var advance int

		if consumed >= len(input) {
			break
		}
		a, advance = nextValidRune(input[consumed:])
		if a == ' ' {
			break
		}
		consumed += advance
		if consumed >= len(input) {
			b = 'x'
		} else {
			b, advance = nextValidRune(input[consumed:])
			if b == ' ' {
				b = 'x'
			}
		}
		if a == b {
			result = append(result, []rune{a, 'x'})
			continue
		}
		result = append(result, []rune{a, b})
		consumed += advance
	}
	return result
}

type moveFunc func(int) int

func (m *Cipher) translate(inputText string, move moveFunc) (string, error) {
	pairs := runePairs(inputText)

	if len(pairs) == 0 {
		return "", errors.New("found no translatable characters in input")
	}

	result := ""

	for _, pair := range pairs {
		var newLocA, newLocB location

		locA := m.locations[pair[0]]
		locB := m.locations[pair[1]]

		newLocA = locA
		newLocB = locB

		if locA.col == locB.col {
			newLocA.row = move(locA.row)
			newLocB.row = move(locB.row)
		} else if locA.row == locB.row {
			newLocA.col = move(locA.col)
			newLocB.col = move(locB.col)
		} else {
			newLocA.col = locB.col
			newLocB.col = locA.col
		}
		newA := m.content[newLocA.row][newLocA.col]
		newB := m.content[newLocB.row][newLocB.col]
		// fmt.Printf("%c%c -> %c%c (%v -> %v) (%v -> %v)\n",
		// 	pair[0], pair[1], newA, newB, locA, newLocA, locB, newLocB)
		result = result + string(newA) + string(newB)
	}

	return result, nil
}

// Encode translates the plain text argument to encrypted text.
func (m *Cipher) Encode(plainText string) (string, error) {
	results, err := m.translate(plainText, next)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode")
	}
	return results, nil
}

func next(i int) int {
	return (i + 1) % 5
}

func prev(i int) int {
	if i == 0 {
		return 4
	}
	return i - 1
}

// Decode translates the encrypted argument to plain text.
func (m *Cipher) Decode(cipherText string) (string, error) {
	results, err := m.translate(cipherText, prev)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode")
	}
	return results, nil
}

// New creates a Cipher using the given keyword, or returns an
// error if the word cannot be used to create a matrix.
func New(keyword string) (*Cipher, error) {
	keyword = strings.ToLower(keyword)
	m := &Cipher{
		keyword:   keyword,
		locations: map[rune]location{},
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
		m.locations[r] = location{row: row, col: col}
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	for _, r := range remaining {
		m.content[row][col] = r
		m.locations[r] = location{row: row, col: col}
		col++
		if col >= 5 {
			col = 0
			row++
		}
	}
	return m, nil
}
