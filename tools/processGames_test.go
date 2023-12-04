package tools

import "testing"

var input = ReadFileToString("../day02/demoPart01.txt")

func TestProcessGames(t *testing.T) {
	type args struct {
		cubeGames string
		part int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 01",
			args: args{cubeGames: input, part: 1},
			want: 8, // Replace 123 with the expected output of the test case
		},
		{
			name: "Part 02",
			args: args{cubeGames: input, part: 2},
			want: 2286, // Replace 123 with the expected output of the test case
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessGames(tt.args.cubeGames, tt.args.part); got != tt.want {
				t.Errorf("ProcessGames() = %v, want %v", got, tt.want)
			}
		})
	}
}
