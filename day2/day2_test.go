package day2

import (
	"fmt"
	"testing"
)

func TestRunIntcode(t *testing.T) {
	var tests = []struct {
		initial string
		want    string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50"},
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.initial)
		t.Run(testname, func(t *testing.T) {
			initial := DeserializeIntcode(tt.initial)
			final := RunIntcode(initial)
			ans := SerializeIntcode(final)
			if ans != tt.want {
				t.Errorf("Got %s, want %s", ans, tt.want)
			}
		})
	}
}
