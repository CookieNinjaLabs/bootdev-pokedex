package pokecache

import (
	"fmt"
	"testing"
	"time"
)

type superContestEffect struct {
	Appeal            int `json:"appeal"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	ID    int `json:"id"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		//{
		//	key: "https://pokeapi.co/api/v2/super-contest-effect/1/",
		//	val: []byte(`{"appeal":2,"flavor_text_entries":[{"flavor_text":"Enables the user to perform first in the next turn.","language":{"name":"en","url":"https://pokeapi.co/api/v2/language/9/"}}],"id":1,"moves":[{"name":"agility","url":"https://pokeapi.co/api/v2/move/97/"},{"name":"quick-attack","url":"https://pokeapi.co/api/v2/move/98/"},{"name":"teleport","url":"https://pokeapi.co/api/v2/move/100/"},{"name":"double-team","url":"https://pokeapi.co/api/v2/move/104/"},{"name":"cotton-spore","url":"https://pokeapi.co/api/v2/move/178/"},{"name":"mach-punch","url":"https://pokeapi.co/api/v2/move/183/"},{"name":"extreme-speed","url":"https://pokeapi.co/api/v2/move/245/"},{"name":"tailwind","url":"https://pokeapi.co/api/v2/move/366/"},{"name":"me-first","url":"https://pokeapi.co/api/v2/move/382/"},{"name":"sucker-punch","url":"https://pokeapi.co/api/v2/move/389/"},{"name":"rock-polish","url":"https://pokeapi.co/api/v2/move/397/"},{"name":"vacuum-wave","url":"https://pokeapi.co/api/v2/move/410/"},{"name":"bullet-punch","url":"https://pokeapi.co/api/v2/move/418/"},{"name":"ice-shard","url":"https://pokeapi.co/api/v2/move/420/"},{"name":"shadow-sneak","url":"https://pokeapi.co/api/v2/move/425/"},{"name":"aqua-jet","url":"https://pokeapi.co/api/v2/move/453/"}]}`),
		//},
		//{
		//	key: "https://pokeapi.co/api/v2/super-contest-effect/2",
		//	val: []byte(`{"appeal":2,"flavor_text_entries":[{"flavor_text":"Enables the user to perform last in the next turn.","language":{"name":"en","url":"https://pokeapi.co/api/v2/language/9/"}}],"id":2,"moves":[{"name":"bubble-beam","url":"https://pokeapi.co/api/v2/move/61/"},{"name":"bubble","url":"https://pokeapi.co/api/v2/move/145/"},{"name":"scary-face","url":"https://pokeapi.co/api/v2/move/184/"},{"name":"icy-wind","url":"https://pokeapi.co/api/v2/move/196/"},{"name":"vital-throw","url":"https://pokeapi.co/api/v2/move/233/"},{"name":"rock-tomb","url":"https://pokeapi.co/api/v2/move/317/"},{"name":"mud-shot","url":"https://pokeapi.co/api/v2/move/341/"},{"name":"hammer-arm","url":"https://pokeapi.co/api/v2/move/359/"}]}`),
		//},
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			got, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to get %v, got %v", c.val, got)
			}
			if string(got) != string(c.val) {
				t.Errorf("Expected to get %v, got %v", c.val, got)
			}
		})
	}
}
