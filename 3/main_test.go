package main

import "testing"

func Test_powerConsumption(t *testing.T) {
	type args struct {
		reports []string
	}
	tests := []struct {
		name string
		args args
		want int64
		oxy  int64
		co2  int64
	}{
		{
			name: "example",
			args: args{
				reports: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			want: 198,
			oxy:  23,
			co2:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerConsumption(tt.args.reports); got != tt.want {
				t.Errorf("powerConsumption() = %v, want %v", got, tt.want)
			}
			if got := findRating(tt.args.reports, byte('1'), byte('0'), true); got != tt.oxy {
				t.Errorf("findRating oxy = %v, want %v", got, tt.oxy)
			}
			if got := findRating(tt.args.reports, byte('0'), byte('1'), false); got != tt.co2 {
				t.Errorf("findRating co2 = %v, want %v", got, tt.co2)
			}
		})
	}
}
