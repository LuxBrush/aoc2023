package tools

import "testing"

var demoPart01 = ReadFileToString("../day03/demoPart01.txt")
var demoPart02 = ReadFileToString("../day03/demoPart02.txt")
var demoPart02P = ReadFileToString("../day03/demoPart02-pattern.txt")

func TestProcessSchematic(t *testing.T) {
	type args struct {
		schematic string
		part      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 01",
			args: args{schematic: demoPart01, part: 1},
			want: 4361,
		},
		{
			name: "Part 02",
			args: args{schematic: demoPart01, part: 2},
			want: 467835,
		},
		{
			name: "Part 02: Large Grid",
			args: args{schematic: demoPart02, part: 2},
			want: 935670,
		},
		{
			name: "Part 02: Pattern",
			args: args{schematic: demoPart02P, part: 2},
			want: 1071447,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessSchematic(tt.args.schematic, tt.args.part); got != tt.want {
				t.Errorf("ProcessSchematic() = %v, want %v", got, tt.want)
			}
		})
	}
}
