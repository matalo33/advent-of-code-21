package main

import (
	"testing"
)

func Test_countLarger(t *testing.T) {
	type args struct {
		measurements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				measurements: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLarger(tt.args.measurements); got != tt.want {
				t.Errorf("countLarger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSlidingWindowsLarger(t *testing.T) {
	type args struct {
		measurements []int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "example",
			args: args{
				measurements: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			wantCount: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countSlidingWindowsLarger(tt.args.measurements); gotCount != tt.wantCount {
				t.Errorf("countSlidingWindowsLarger() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
