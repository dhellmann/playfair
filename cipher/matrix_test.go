package cipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix(t *testing.T) {
	for _, tc := range []struct {
		Keyword  string
		Expected string
	}{
		{
			Keyword: "playfairexample",
			Expected: `p l a y f 
i r e x m 
b c d g h 
k n o q s 
t u v w z 
`,
		},
		{
			Keyword: "playfajrexample",
			Expected: `p l a y f 
i r e x m 
b c d g h 
k n o q s 
t u v w z 
`,
		},
		{
			Keyword: "monarchy",
			Expected: `m o n a r 
c h y b d 
e f g i k 
l p q s t 
u v w x z 
`,
		},
		{
			Keyword: "MONARCHY",
			Expected: `m o n a r 
c h y b d 
e f g i k 
l p q s t 
u v w x z 
`,
		},
	} {
		t.Run(tc.Keyword, func(t *testing.T) {
			m, _ := New(tc.Keyword)
			s := m.Matrix()
			assert.Equal(t, tc.Expected, s)
		})
	}
}

func TestNewNonASCII(t *testing.T) {
	for _, candidate := range []string{
		"with a space",
		"with-punctuation",
		"withumlaüt",
		"short",
	} {
		t.Run(candidate, func(t *testing.T) {
			_, err := New(candidate)
			assert.Error(t, err)
		})
	}
}

func TestEncode(t *testing.T) {
	for _, tc := range []struct {
		Keyword  string
		Input    string
		Expected string
	}{
		{
			Keyword:  "playfairexample",
			Input:    "Hide the gold in the tree stump",
			Expected: "bmodzbxdnabekudmuixmmouvif",
		},
		{
			Keyword:  "monarchy",
			Input:    "instruments",
			Expected: "gatlmzclrqxa",
		},
	} {
		t.Run(tc.Keyword, func(t *testing.T) {
			m, err := New(tc.Keyword)
			assert.Nil(t, err)
			actual, err := m.Encode(tc.Input)
			assert.Nil(t, err)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}

func TestEncodeError(t *testing.T) {
	for _, tc := range []string{
		"",
		" ",
		"\n",
	} {
		t.Run(tc, func(t *testing.T) {
			m, _ := New("playfair")
			_, err := m.Encode(tc)
			assert.Error(t, err)
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []struct {
		Keyword  string
		Input    string
		Expected string
	}{
		{
			Keyword:  "playfairexample",
			Input:    "bmodzbxdnabekudmuixmmouvif",
			Expected: "hidethegoldinthetrexestump",
		},
		{
			Keyword:  "playfairexample",
			Input:    "bm od zb xd na be ku dm ui xm mo uv if",
			Expected: "hidethegoldinthetrexestump",
		},
		{
			Keyword:  "monarchy",
			Input:    "gatlmzclrqxa",
			Expected: "instrumentsx",
		},
	} {
		t.Run(tc.Keyword, func(t *testing.T) {
			m, err := New(tc.Keyword)
			assert.Nil(t, err)
			actual, err := m.Decode(tc.Input)
			assert.Nil(t, err)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}

func TestDecodeError(t *testing.T) {
	for _, tc := range []string{
		"",
		" ",
		"\n",
	} {
		t.Run(tc, func(t *testing.T) {
			m, _ := New("playfair")
			_, err := m.Decode(tc)
			assert.Error(t, err)
		})
	}
}

func TestNextValidRune(t *testing.T) {
	for _, tc := range []struct {
		Input    string
		Expected rune
		Consumed int
	}{
		{
			Input:    "abc",
			Expected: 'a',
			Consumed: 1,
		},
		{
			Input:    " abc",
			Expected: 'a',
			Consumed: 2,
		},
		{
			Input:    "",
			Expected: ' ',
			Consumed: 0,
		},
		{
			Input:    "j",
			Expected: 'i',
			Consumed: 1,
		},
	} {
		t.Run(tc.Input, func(t *testing.T) {
			r, c := nextValidRune(tc.Input)
			assert.Equal(t, tc.Expected, r)
			assert.Equal(t, tc.Consumed, c)
		})
	}
}

func TestRunePairs(t *testing.T) {
	for _, tc := range []struct {
		Input    string
		Expected [][]rune
	}{
		{
			Input:    "a",
			Expected: [][]rune{{'a', 'x'}},
		},
		{
			Input:    "a-",
			Expected: [][]rune{{'a', 'x'}},
		},
		{
			Input:    "abc",
			Expected: [][]rune{{'a', 'b'}, {'c', 'x'}},
		},
		{
			Input:    " abc",
			Expected: [][]rune{{'a', 'b'}, {'c', 'x'}},
		},
		{
			Input:    "aabc",
			Expected: [][]rune{{'a', 'x'}, {'a', 'b'}, {'c', 'x'}},
		},
	} {
		t.Run(tc.Input, func(t *testing.T) {
			pairs := runePairs(tc.Input)
			assert.Equal(t, tc.Expected, pairs)
		})
	}
}
