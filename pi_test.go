package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPiDigits(t *testing.T) {
	type TestCase struct {
		name string
		lower int
		upper int
		expected string
	}

	tests := []TestCase{
		{
			name: "First 5 of all digits",
			lower: 0,
			upper: 5,
			expected: "3.141",
		},
		{
			name: "First 5 decimal places",
			lower: 2,
			upper: 7,
			expected: "14159",
		},
		{
			name: "45th to 50th DP",
			lower: 47,
			upper: 52,
			expected: "37510",
		},
		{
			name: "95th to 100th DP",
			lower: 97,
			upper: 102,
			expected: "70679",
		},
		{
			name: "5000th to 5005th DP",
			lower: 5002,
			upper: 5007,
			expected: "56951",
		},
		{
			name: "Last five",
			lower: len(PI)-5,
			upper: len(PI),
			expected: "62541",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, PI[tc.lower:tc.upper], "digits don't match")
		})
	}
}
