package main

import "testing"

func Test_findHorizontalPosition(t *testing.T) {
	type args struct {
		crabs []int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantp2 int
	}{
		{
			name: "example",
			args: args{
				crabs: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			},
			want:   37,
			wantp2: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findHorizontalPosition(tt.args.crabs, true); got != tt.want {
				t.Errorf("findHorizontalPosition() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findHorizontalPosition(tt.args.crabs, false); got != tt.wantp2 {
				t.Errorf("findHorizontalPosition() Part 2 = %v, want %v", got, tt.wantp2)
			}
		})
	}
}
