package day4

import (
	"fmt"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	var tests = []struct {
		initial string
		want    bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.initial)
		t.Run(testname, func(t *testing.T) {
			ans := IsValidPassword(tt.initial)
			if ans != tt.want {
				t.Errorf("Got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestIsValidPassword2(t *testing.T) {
	var tests = []struct {
		initial string
		want    bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.initial)
		t.Run(testname, func(t *testing.T) {
			ans := IsValidPassword2(tt.initial)
			if ans != tt.want {
				t.Errorf("Got %t, want %t, for %s", ans, tt.want, tt.initial)
			}
		})
	}
}
