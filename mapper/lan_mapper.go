package mapper

import "fmt"

type lanMapper struct {
	s string
	p int
}

func (m lanMapper) TransformRune(pos int) {
	if r, e := upperEveryNChars(pos, m.s); e != nil {
		fmt.Println(fmt.Sprintf("problems with %v pos %v err [%v]", m.s, pos, e))
	} else {
		fmt.Println(fmt.Sprintf("pos %v %v", pos, string(r)))
	}
}

func (m lanMapper) GetValueAsRuneSlice() []rune {
	return []rune(m.s)
}

func (m lanMapper) String() string {
	r, _ := upperEveryNChars(m.p, m.s)
	return fmt.Sprintf("Lan Mapper mapped string %v og string %v pos %v", string(r), m.s, m.p)
}
