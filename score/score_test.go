package score

import (
	"testing"

	"github.com/ascii-arcade/knucklebones/dice"
)

func TestCalculate(t *testing.T) {
	type args struct {
		pool []dice.DicePool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty Pool",
			args: args{pool: []dice.DicePool{}},
			want: 0,
		},
		{
			name: "Single Column with Single Die",
			args: args{pool: []dice.DicePool{{1}}},
			want: 1,
		},
		{
			name: "Single Column with Multiple Dice",
			args: args{pool: []dice.DicePool{{1, 2, 3}}},
			want: 6,
		},
		{
			name: "Multiple Columns with Different Faces",
			args: args{pool: []dice.DicePool{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}},
			want: (1 + 2 + 3) * 3,
		},
		{
			name: "Single Column with Duplicates",
			args: args{pool: []dice.DicePool{
				{1, 1},
			}},
			want: 4,
		},
		{
			name: "Single Column with Triples",
			args: args{pool: []dice.DicePool{
				{2, 2, 2},
			}},
			want: 18,
		},
		{
			name: "Multiple Columns with Duplicates",
			args: args{pool: []dice.DicePool{
				{1, 1},
				{2, 2},
			}},
			want: 12,
		},
		{
			name: "Mix",
			args: args{pool: []dice.DicePool{
				{1, 2, 3},
				{2, 2, 3},
				{6, 6, 6},
			}},
			want: 71,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.pool); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
