package model

import "testing"

func TestParsePriorit(t *testing.T) {
	tests := []struct {
		input    string
		expected Priority
	}{
		{"low", Low},
		{"medium", Medium},
		{"high", High},
		{"unknown", Low},
	}

	for _, tt := range tests {
		got := ParsePriority(tt.input)

		if *got != tt.expected {
			t.Errorf("Parse priority(%s): expected %v, got %v", tt.input, tt.expected, *got)
		}
	}
}
