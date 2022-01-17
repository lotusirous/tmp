package isbn

import "testing"

func TestIsISBN(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{
			arg:  "1112223339",
			want: true,
		},
		{
			arg:  "111222333",
			want: false,
		},
		{
			arg:  "1112223339X",
			want: false,
		},
		{
			arg:  "1234554321",
			want: true,
		},
		{
			arg:  "1234512345",
			want: false,
		},
		{
			arg:  "048665088X",
			want: true,
		},
		{
			arg:  "X123456788",
			want: false,
		},
		{
			arg:  "1293",
			want: false,
		},
		{
			arg:  "048665088x",
			want: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.arg, func(t *testing.T) {
			if got := ValidISBN10(tc.arg); got != tc.want {
				t.Errorf("ValidISBN10(%s) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}
