// Copyright 2020 Lan Nguyen. All rights reserved.
package main

import (
	"fmt"

	"github.com/lanny1406/aspiration/mapper"
)

//
// Capitalizes *only* every third word character of a string
// Characters are considered letters from "Western Alphabet" [a-z A-Z]

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(s)
	fmt.Println(s)
}

// NewSkipString capitalizes every N character for a given string
func NewSkipString(pos int, s string) mapper.Interface {
	return mapper.NewSkipString(pos, s)
}
