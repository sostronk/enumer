// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Check that constants defined as a conversion are accepted.

package main

import (
	"fmt"
	"strings"
)

type Other int // Imagine this is in another package.

const (
	alpha Other = iota
	beta
	gamma
	delta
)

type Conv int

const (
	Alpha = Conv(alpha)
	Beta  = Conv(beta)
	Gamma = Conv(gamma)
	Delta = Conv(delta)
)

func main() {
	ck(Alpha, "Alpha")
	ck(Beta, "Beta")
	ck(Gamma, "Gamma")
	ck(Delta, "Delta")
	ck(42, "Conv(42)")
}

func ck(c Conv, str string) {
	if fmt.Sprint(c) != str {
		panic("conv.go: " + str)
	}

	v, ok := ParseConv(str)
	if strings.HasPrefix(str, "Conv(") {
		if ok {
			panic("conv.go: should not parse " + str)
		}
		return
	}
	if !ok {
		panic("conv.go: could not parse " + str)
	}
	if v != c {
		panic("conv.go: wrong parse " + str)
	}
}
