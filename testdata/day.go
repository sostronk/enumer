// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple test: enumeration of type int starting at 0.

package main

import (
	"fmt"
	"strings"
)

type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
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

	var v Day
	err := v.UnmarshalText([]byte(str))
	if strings.HasPrefix(str, "Day(") {
		if err == nil {
			panic("day.go: should not parse " + str)
		}
		return
	}
	if err != nil {
		panic("day.go: could not parse " + str)
	}
	if v != day {
		panic("day.go: wrong parse " + str)
	}
}
