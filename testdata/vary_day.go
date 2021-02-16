// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is the same as day.go except the constants have different values.

package main

import (
	"fmt"
	"strings"
)

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	ck(Monday, "Monday")
	ck(Tuesday, "Tuesday")
	ck(Wednesday, "Wednesday")
	ck(Thursday, "Thursday")
	ck(Friday, "Friday")
	ck(Saturday, "Saturday")
	ck(Sunday, "Sunday")
	ck(-127, "Day(-127)")
	ck(127, "Day(127)")
}

func ck(day Day, str string) {
	if fmt.Sprint(day) != str {
		panic("day.go: " + str)
	}

	v, ok := ParseDay(str)
	if strings.HasPrefix(str, "Day(") {
		if ok {
			panic("day.go: should not parse " + str)
		}
		return
	}
	if !ok {
		panic("day.go: could not parse " + str)
	}
	if v != day {
		panic("day.go: wrong parse " + str)
	}
}
