// https://leetcode.com/problems/number-of-recent-calls/
package main

import "fmt"

type RecentCounter struct {
	calls []int
	idx   int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: make([]int, 0, 10000),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)

	tr := t - 3000

	for len(this.calls)-this.idx > 0 && this.calls[this.idx] < tr {
		this.idx++
	}

	return len(this.calls) - this.idx
}

func main() {
	r := Constructor()

	fmt.Println(r.Ping(1))
	fmt.Println(r.Ping(100))
	fmt.Println(r.Ping(3001))
	fmt.Println(r.Ping(3002))
}
