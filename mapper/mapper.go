package mapper

import (
	"fmt"
	"strings"
	"unicode"
)

// Interface uppercases (n*pos) of a string
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// MapString ...
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// NewSkipString returns a new mapper interface
func NewSkipString(pos int, s string) Interface {
	return &lanMapper{s: s, p: pos}
}

// Every3rdChar converts every third char to uppercase
func Every3rdChar(s string) string {
	n := 3
	r, err := upperEveryNChars(n, s)
	if err != nil {
		fmt.Println(fmt.Sprintf("problems with %v n %v err [%v]", s, n, err))
	}
	return string(r)
}

// lowercases string before deciding which char to uppercase
func upperEveryNChars(n int, s string) ([]rune, error) {
	s = strings.ToLower(s)
	if n <= 0 {
		return []rune(s), fmt.Errorf("invalid index")
	}

	if len(s) < n {
		return []rune(s), nil
	}

	ss := ""
	upperLimit := n - 1
	for _, c := range s {
		cc := string(c)
		if !unicode.IsLetter(c) {
			// not a letter
			ss = ss + cc
			continue
		}
		if upperLimit == 0 {
			cc = strings.ToUpper(cc)
			upperLimit = n
		}
		upperLimit = upperLimit - 1
		ss = ss + cc
	}
	return []rune(ss), nil
}
