package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input:   "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual does not equal expected")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if word != expectedWord {
				t.Errorf("%s (actual) != %s (expected)", word, expectedWord)
				continue
			}
		}
	}

}