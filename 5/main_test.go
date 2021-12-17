package main

import "testing"

func Test_overlappingPoints(t *testing.T) {
	type args struct {
		lines    []line
		gridSize int
	}
	tests := []struct {
		name         string
		args         args
		want         int
		wantDiagonal int
	}{
		{
			name: "example",
			args: args{
				lines: []line{
					{0, 9, 5, 9},
					{8, 0, 0, 8},
					{9, 4, 3, 4},
					{2, 2, 2, 1},
					{7, 0, 7, 4},
					{6, 4, 2, 0},
					{0, 9, 2, 9},
					{3, 4, 1, 4},
					{0, 0, 8, 8},
					{5, 5, 8, 2},
				},
				gridSize: 10,
			},
			want:         5,
			wantDiagonal: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overlappingPoints(tt.args.lines, tt.args.gridSize, false); got != tt.want {
				t.Errorf("overlappingPoints Straight = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := overlappingPoints(tt.args.lines, tt.args.gridSize, true); got != tt.wantDiagonal {
				t.Errorf("overlappingPoints Diagonal = %v, want %v", got, tt.wantDiagonal)
			}
		})
	}
}
