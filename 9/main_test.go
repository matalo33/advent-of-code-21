package main

import "testing"

func Test_getRiskLevel(t *testing.T) {
	type args struct {
		heightmap [][]int
	}
	tests := []struct {
		name                  string
		args                  args
		wantLowPointRiskLevel int
		wantBasinRiskLevel    int
	}{
		{
			name: "example",
			args: args{
				heightmap: [][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
			},
			wantLowPointRiskLevel: 15,
			wantBasinRiskLevel:    1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lowPoints := getLowPoints(tt.args.heightmap)
			if gotRiskLevel := getRiskLevel(lowPoints, tt.args.heightmap); gotRiskLevel != tt.wantLowPointRiskLevel {
				t.Errorf("getRiskLevel() = %v, want %v", gotRiskLevel, tt.wantLowPointRiskLevel)
			}
			if gotRiskLevel := getBasinsRiskLevel(lowPoints, tt.args.heightmap); gotRiskLevel != tt.wantBasinRiskLevel {
				t.Errorf("getBasinsRiskLevel() = %v, want %v", gotRiskLevel, tt.wantBasinRiskLevel)
			}
		})
	}
}
