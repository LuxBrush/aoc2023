package tools

import "testing"

var inputDemoDay03 = ReadFileToString("../day03/demoPart01.txt")

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
			args: args{schematic: inputDemoDay03, part: 1},
			want: 4361,
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
