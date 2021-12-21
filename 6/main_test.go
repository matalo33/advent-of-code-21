package main

import "testing"

func Test_makeFish(t *testing.T) {
	type args struct {
		fish []int
		days int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "80",
			args: args{
				fish: []int{3, 4, 3, 1, 2},
				days: 80,
			},
			want: 5934,
		},
		{
			name: "256",
			args: args{
				fish: []int{3, 4, 3, 1, 2},
				days: 256,
			},
			want: 26984457539,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFishV2(tt.args.fish, tt.args.days); got != tt.want {
				t.Errorf("makeFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
