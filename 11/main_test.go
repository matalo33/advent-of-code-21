package main

import "testing"

func Test_countFlashes(t *testing.T) {
	type args struct {
		octopi [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{octopi: [][]int{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			}},
			want: 1656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFlashes(tt.args.octopi, 100); got != tt.want {
				t.Errorf("countFlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
