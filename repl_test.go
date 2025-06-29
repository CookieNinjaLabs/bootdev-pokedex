package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		got  string
		want []string
	}{
		{
			got:  "hi my name is",
			want: []string{"hi", "my", "name", "is"},
		},
		{
			got:  "",
			want: []string{},
		},
		{
			got:  "  hello world  ",
			want: []string{"hello", "world"},
		},
		{
			got:  "Hello World",
			want: []string{"hello", "world"},
		},
	}

	// Loop through all test cases and input the gots and wants
	for _, c := range cases {
		got := cleanInput(c.got)
		want := c.want

		// Check if the slices have the same length otherwise return an error
		if len(got) != len(want) {
			t.Errorf("Expected %d, got %d", len(want), len(got))
		}

		// loop through all items of the got and want slice if they have the same
		// length and compare the strings
		// if they don't match return an error
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("Expected %s, got %s", want[i], got[i])
			}
		}
	}
}
