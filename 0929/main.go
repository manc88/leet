package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{
		"test.email+alex@leetcode.com", "test.email.leet+alex@code.com",
	}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{}, len(emails))

	for _, v := range emails {
		m[sanitize(v)] = struct{}{}
	}

	fmt.Println(m)

	return len(m)
}

func sanitize(email string) string {
	s := strings.Split(email, "@")
	local := s[0]
	localbeforePlus := strings.Split(local, "+")[0]

	final := strings.Replace(localbeforePlus, ".", "", -1) + "@" + s[1]
	fmt.Println(final)
	return final
}
