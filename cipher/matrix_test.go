package cipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix(t *testing.T) {
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
	} {
		t.Run(tc.Keyword, func(t *testing.T) {
			m, _ := NewMatrix(tc.Keyword)
			s := m.String()
			assert.Equal(t, tc.Expected, s)

		})
	}
}
