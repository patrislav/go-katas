package wordchain

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func ExampleBuild() {
	r := strings.NewReader("cat dog dag dig cot cut cad cog joy apple")
	dict := NewDictionary(3)
	_, _ = dict.ReadFrom(r)

	chain, _ := Build(dict, "cat", "dog")
	fmt.Println(chain)
	// Output: [cat cot cog dog]
}

func TestBuild(t *testing.T) {
	catDog := []string{"cat", "dog", "dag", "dig", "dug", "cot", "cut", "cad", "cog", "joy", "apple"}
	rubyCode := []string{"robs", "ruby", "code", "core", "code", "cozy", "rods", "rubs", "rode"}
	tests := []struct {
		initial, final string
		words          []string
		result         []string
		err            error
	}{
		{"", "", catDog, nil, ErrEmpty},
		{"cat", "apple", catDog, nil, ErrNotEqualLength},
		{"cat", "joy", catDog, nil, ErrNotFound},
		{"pig", "dog", catDog, nil, ErrNotInDictionary},
		{"cat", "dog", catDog, []string{"cat", "cot", "cog", "dog"}, nil},
		{"ruby", "code", rubyCode, []string{"ruby", "rubs", "robs", "rods", "rode", "code"}, nil},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s to %s", tc.initial, tc.final), func(t *testing.T) {
			dict := NewDictionary(len(tc.initial))
			for _, w := range tc.words {
				dict.AddWord(w)
			}
			res, err := Build(dict, tc.initial, tc.final)
			if err != tc.err {
				t.Errorf("error: got %v, want %v", err, tc.err)
			}
			if !reflect.DeepEqual(res, tc.result) {
				t.Errorf("result: got %v, want %v", res, tc.result)
			}
		})
	}

	// pairs with multiple possible solutions (paths) are unstable - the chain might be different each iteration
	t.Run("bear to fish", func(t *testing.T) {
		words := []string{
			"bear", "fish", "beat", "best", "fiat", "fist", "fest", "feat", "fear", "beak", "biak", "bisk", "bish",
		}
		dict := NewDictionary(4)
		for _, w := range words {
			dict.AddWord(w)
		}
		res, err := Build(dict, "bear", "fish")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(res) != 6 {
			t.Errorf("expected a chain of 6 words, got %v", len(res))
		}
	})
}

func BenchmarkBuild(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/wordlist.txt")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}
	wordlist := string(buf)
	for n := 0; n < b.N; n++ {
		r := strings.NewReader(wordlist)
		dict := NewDictionary(4)
		_, err := dict.ReadFrom(r)
		if err != nil {
			b.Errorf("unexpected error: %v", err)
		}
		_, err = Build(dict, "ruby", "code")
		if err != nil {
			b.Errorf("unexpected error: %v", err)
		}
	}
}
