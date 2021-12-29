package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		part2     int
	}{
		{
			name: "example",
			args: args{
				lines: []string{
					"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
					"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
					"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
					"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
					"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
					"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
					"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
					"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
					"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
					"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
				},
			},
			wantCount: 26,
			part2:     61229,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := part1(tt.args.lines); gotCount != tt.wantCount {
				t.Errorf("part1() = %v, want %v", gotCount, tt.wantCount)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := part2(tt.args.lines); gotCount != tt.part2 {
				t.Errorf("part2() = %v, want %v", gotCount, tt.part2)
			}
		})
	}
}

func Test_leftNotInRight(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "3",
			args: args{
				left:  "cf",
				right: "abcdfg",
			},
			wantCount: 0,
		},
		{
			name: "2",
			args: args{
				left:  "cf",
				right: "acdeg",
			},
			wantCount: 1,
		},
		{
			name: "5",
			args: args{
				left:  "cf",
				right: "abdfg",
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := leftNotInRight(tt.args.left, tt.args.right); len(gotCount) != tt.wantCount {
				t.Errorf("leftNotInRight() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
