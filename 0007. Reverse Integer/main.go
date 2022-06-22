//https://leetcode.com/problems/reverse-integer/
package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {

	xstring := strconv.Itoa(x)

	negative := xstring[0] == '-'

	var rns []rune
	if negative {
		rns = []rune(xstring[1:])
	} else {
		rns = []rune(xstring)
	}

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	revstring := string(rns)
	if negative {
		revstring = "-" + revstring
	}

	var revint int
	var err error
	if revint, err = strconv.Atoi(revstring); err != nil || revint > math.MaxInt32 || revint < math.MaxInt32*-1 {
		return 0
	}

	return revint

}
