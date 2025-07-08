package pokeapi

import (
	"fmt"
	"strings"
	"testing"
)

func TestCatchPokemonResultMessage(t *testing.T) {
	cases := []struct {
		id      int
		got     string
		want    string
		err     error
		success bool
	}{
		{
			id:      1,
			got:     "Elektek",
			want:    "Elektek was caught",
			err:     nil,
			success: true,
		},
		{
			id:      2,
			got:     "Pikachu",
			want:    "Pikachu was caught",
			err:     nil,
			success: true,
		},
		{
			id:      3,
			got:     "",
			want:    "",
			err:     fmt.Errorf("pokemon is empty"),
			success: true,
		},
		{
			id:      4,
			got:     "Elektek",
			want:    "Elektek escaped",
			err:     nil,
			success: false,
		},
		{
			id:      5,
			got:     "Pikachu",
			want:    "Pikachu escaped",
			err:     nil,
			success: false,
		},
	}
	for _, c := range cases {
		got, err := catchPokemonResultMessage(c.got, c.success)
		want := c.want

		if err != nil && c.err == nil {
			t.Errorf("got error on case id %d: %s", c.id, err)
		}

		if c.err != nil {
			if err == nil {
				t.Errorf("Expected error %v, got nil", c.err)
			}

			if !strings.Contains(err.Error(), c.err.Error()) {
				t.Errorf("Expected error %v, got %v", c.err, err)
			}
		}
		if got != want {
			t.Errorf("Expected %s, got %s", want, got)
		}
	}
}
