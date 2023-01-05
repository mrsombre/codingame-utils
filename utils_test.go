package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected string
	}{
		{
			name:     `>0`,
			x:        1,
			expected: "1",
		},
		{
			name:     `<0`,
			x:        -1,
			expected: "-1",
		},
		{
			name:     `0`,
			x:        0,
			expected: "0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, intToStr(tc.x))
		})
	}
}

func TestBoolToInt(t *testing.T) {
	tests := []struct {
		name     string
		x        bool
		expected int
	}{
		{
			name:     `true`,
			x:        true,
			expected: 1,
		},
		{
			name:     `false`,
			x:        false,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, boolToInt(tc.x))
		})
	}
}
