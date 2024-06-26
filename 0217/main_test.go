// https://leetcode.com/problems/contains-duplicate/description/
package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				nums: []int{},
			},
			want: false,
		},
		{
			name: "one",
			args: args{
				nums: []int{1},
			},
			want: false,
		},
		{
			name: "only 1 digit",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: true,
		},
		{
			name: "no dupl",
			args: args{
				nums: []int{4, 1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
