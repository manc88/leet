package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Expected: MMCMXCIV")
	intToRoman(2994)

}

func intToRoman(num int) string {

	res := ""

	ms := num / 1000
	ds := (num - (ms * 1000)) / 500
	cs := (num - (ms * 1000) - (ds * 500)) / 100
	ls := (num - (ms * 1000) - (ds * 500) - (cs * 100)) / 50
	xs := (num - (ms * 1000) - (ds * 500) - (cs * 100) - (ls * 50)) / 10
	vs := (num - (ms * 1000) - (ds * 500) - (cs * 100) - (ls * 50) - (xs * 10)) / 5
	is := (num - (ms * 1000) - (ds * 500) - (cs * 100) - (ls * 50) - (xs * 10) - (vs * 5))

	//fmt.Println("m=", ms, "d=", ds, "c=", cs, "l=", ls, "x=", xs, "v=", vs, "i=", is)

	for i := 0; i < ms; i++ {
		res += "M"
	}
	for i := 0; i < ds; i++ {
		res += "D"
	}
	for i := 0; i < cs; i++ {
		res += "C"
	}
	for i := 0; i < ls; i++ {
		res += "L"
	}
	for i := 0; i < xs; i++ {
		res += "X"
	}
	for i := 0; i < vs; i++ {
		res += "V"
	}
	for i := 0; i < is; i++ {
		res += "I"
	}

	replacements := []string{
		"DCCCC:CM",
		"CCCC:CD",
		"LXXXX:XC",
		"XXXX:XL",
		"VIIII:IX",
		"IIII:IV",
	}
	for _, s := range replacements {
		sp := strings.Split(s, ":")
		res = strings.Replace(res, sp[0], sp[1], -1)
	}

	//fmt.Println(num, "=", res)
	//fmt.Println("-------------------")
	return res
}
