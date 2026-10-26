package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"No 'x' in Nixon", true},
		{"0P", false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
