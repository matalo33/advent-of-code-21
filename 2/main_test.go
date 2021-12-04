package main

import "testing"

func Test_calcFinalPosition(t *testing.T) {
	type args struct {
		instructions []instruction
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantWithAim int
	}{
		{
			name: "example",
			args: args{
				instructions: []instruction{
					{
						direction: "forward",
						amount:    5,
					},
					{
						direction: "down",
						amount:    5,
					},
					{
						direction: "forward",
						amount:    8,
					},
					{
						direction: "up",
						amount:    3,
					},
					{
						direction: "down",
						amount:    8,
					},
					{
						direction: "forward",
						amount:    2,
					},
				},
			},
			want:        150,
			wantWithAim: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFinalPosition(tt.args.instructions); got != tt.want {
				t.Errorf("calcFinalPosition() = %v, want %v", got, tt.want)
			}
			if got := calcFinalPositionWithAim(tt.args.instructions); got != tt.wantWithAim {
				t.Errorf("calcFinalPositionWithAim() = %v, want %v", got, tt.wantWithAim)
			}
		})
	}
}
