//https://leetcode.com/problems/string-to-integer-atoi/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//fmt.Printf("%d\n", '.')
	fmt.Println(myAtoi(" -76 fgh"))

}

func myAtoi(s string) int {

	var tmp []rune
	sg := ""
	stop := false
	for _, v := range strings.TrimLeftFunc(s, func(r rune) bool { return r == ' ' }) {

		if stop {
			break
		}

		switch {
		case v == ' ' || v == '.':
			stop = true
		case v == 43 && sg == "":
			sg = "+"
		case v == 43 && sg != "":
			stop = true
		case v == 45 && sg == "":
			sg = "-"
		case v == 45 && sg != "":
			stop = true
		case v >= 48 && v <= 57:
			tmp = append(tmp, v)
			if sg == "" {
				sg = "+"
			}
		default:
			stop = true
		}

	}

	if len(tmp) == 0 {
		return 0
	}

	v, err := strconv.Atoi(sg + string(tmp))

	if err == strconv.ErrRange || v > 2147483647 || v < -2147483648 {
		if sg == "-" {
			return -2147483648
		}
		return 2147483647
	}

	return v
}
