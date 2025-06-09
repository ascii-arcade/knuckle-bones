package score

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		dieFaces  []int
		want      int
		wantError string
	}{
		{
			dieFaces:  []int{1, 1, 1, 4, 5, 5, 6},
			want:      0,
			wantError: "too many dice",
		},
		{
			dieFaces:  []int{},
			want:      0,
			wantError: "no dice",
		},
		{
			dieFaces:  []int{0, 1},
			want:      0,
			wantError: "invalid die face: 0",
		},
		{
			dieFaces:  []int{1, 7},
			want:      0,
			wantError: "invalid die face: 7",
		},
		{
			dieFaces:  []int{1, 1, 1, 1, 2, 5},
			want:      0,
			wantError: "useless dice detected",
		},
		{
			dieFaces: []int{4, 4, 4, 4, 4, 4},
			want:     3000,
		},
		{
			dieFaces: []int{5, 5, 5, 5, 5, 5},
			want:     3000,
		},
		{
			dieFaces: []int{1, 1, 1, 1, 1, 5},
			want:     2050,
		},
		{
			dieFaces: []int{1, 1, 1, 1, 5},
			want:     1050,
		},
		{
			dieFaces: []int{3, 3, 3, 5, 5, 5},
			want:     2500,
		},
		{
			dieFaces: []int{3, 3, 3, 3, 5, 5},
			want:     1500,
		},
		{
			dieFaces: []int{1, 1, 3, 3, 4, 4},
			want:     1500,
		},
		{
			dieFaces: []int{1, 2, 3, 4, 5, 6},
			want:     1500,
		},
		{
			dieFaces: []int{4, 4, 4, 4, 4, 5},
			want:     2050,
		},
		{
			dieFaces: []int{4, 4, 4, 4, 4},
			want:     2000,
		},
		{
			dieFaces: []int{5, 4, 4, 4, 4, 4},
			want:     2050,
		},
		{
			dieFaces: []int{1, 1, 1, 1, 1},
			want:     2000,
		},
		{
			dieFaces: []int{1, 1, 1, 1},
			want:     1000,
		},
		{
			dieFaces: []int{1, 1, 1, 1, 5},
			want:     1050,
		},
		{
			dieFaces: []int{1, 1, 1, 5},
			want:     350,
		},
		{
			dieFaces: []int{5, 5, 5},
			want:     500,
		},
		{
			dieFaces: []int{3, 3, 3, 5},
			want:     350,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Calculate(tt.dieFaces, false)

			if tt.wantError != "" {
				if err == nil {
					t.Errorf("Calculate(%v) = no error, want %v", tt.dieFaces, tt.wantError)
				} else if err.Error() != tt.wantError {
					t.Errorf("Calculate(%v) = %v, want %v", tt.dieFaces, err.Error(), tt.wantError)
				}
			}

			if got != tt.want {
				t.Errorf("Calculate(%v) = %d, want %d", tt.dieFaces, got, tt.want)
			}
		})
	}
}
