// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unsigned integers - check maximum size

package main

import (
	"fmt"
	"strings"
)

type Unum2 uint8

const (
	Zero Unum2 = iota
	One
	Two
)

func main() {
	ck(Zero, "Zero")
	ck(One, "One")
	ck(Two, "Two")
	ck(3, "Unum2(3)")
	ck(255, "Unum2(255)")
}

func ck(unum Unum2, str string) {
	if fmt.Sprint(unum) != str {
		panic("unum.go: " + str)
	}

	var v Unum2
	err := v.UnmarshalText([]byte(str))
	if strings.HasPrefix(str, "Unum2(") {
		if err == nil {
			panic("unum.go: should not parse " + str)
		}
		return
	}
	if err != nil {
		panic("unum.go: could not parse " + str)
	}
	if v != unum {
		panic("unum.go: wrong parse " + str)
	}
}
