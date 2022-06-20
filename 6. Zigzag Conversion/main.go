//https://leetcode.com/problems/zigzag-conversion/
package main

import "strings"

func convert(s string, numRows int) string {

	rows := make([]string, len(s))

	rowptr, direction := 0, 0

	for _, r := range s {

		rows[rowptr] += string(r)

		if rowptr == 0 || rowptr == numRows-1 {

			if direction == 0 {
				direction = 1
			} else {
				direction = 0
			}
		}

		if direction == 0 {
			rowptr--
		} else {
			rowptr++
		}

	}

	return strings.Join(rows, "")

}
