package main

import (
	"reflect"
	"testing"
)

func TestBankDraft(t *testing.T) {
	tests := []struct {
		about    string
		input    int
		expected map[string]int
	}{
		{
			about:    "Test bank draft: 30",
			input:    30,
			expected: map[string]int{"100": 0, "50": 0, "20": 1, "10": 1},
		},
		{
			about:    "Test bank draft: 80",
			input:    80,
			expected: map[string]int{"100": 0, "50": 1, "20": 1, "10": 1},
		},
		{
			about:    "Test bank draft: 15",
			input:    15,
			expected: map[string]int{},
		},
		{
			about:    "Test bank draft: -15",
			input:    -15,
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			res, _ := BankDraft(test.input)
			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, res)
			}
		})
	}
}
