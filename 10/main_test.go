package main

import "testing"

func Test_errorScore(t *testing.T) {
	type args struct {
		navs []string
	}
	tests := []struct {
		name                string
		args                args
		wantErrorScore      int
		wantCompletionScore int
	}{
		{
			name: "example",
			args: args{
				navs: []string{
					"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]",
				},
			},
			wantErrorScore:      26397,
			wantCompletionScore: 288957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorScore, completionScore := getScores(tt.args.navs)
			if got := errorScore; got != tt.wantErrorScore {
				t.Errorf("error score = %v, want %v", got, tt.wantErrorScore)
			}
			if got := completionScore; got != tt.wantCompletionScore {
				t.Errorf("completion score = %v, want %v", got, tt.wantCompletionScore)
			}
		})
	}
}
