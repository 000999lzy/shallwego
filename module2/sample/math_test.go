package sample

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"2x3", 2, 3, 6},
		{"0x5", 0, 5, 0},
		{"-2x4", -2, 4, -8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Multiply(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
