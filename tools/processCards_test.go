package tools

import "testing"

var testPart01 = ReadFileToString("../day04/testPart01.txt")

func TestProcessCard(t *testing.T) {
	type args struct {
		rawCardList string
		part        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 01",
			args: args{rawCardList: testPart01, part: 1},
			want: 13,
		},
		{
			name: "Part 02",
			args: args{rawCardList: testPart01, part: 2},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessCard(tt.args.rawCardList, tt.args.part); got != tt.want {
				t.Errorf("ProcessCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
