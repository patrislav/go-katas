package wordchain

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func ExampleDictionary() {
	words := []string{"dog", "cat", "cog", "cot", "apple"}
	r := strings.NewReader(strings.Join(words, " "))

	dict := NewDictionary(3)
	_, _ = dict.ReadFrom(r)

	fmt.Println("dog", dict.HasWord("dog"))
	fmt.Println("apple", dict.HasWord("apple"))
	// Output:
	// dog true
	// apple false
}

func ExampleDictionary_alphabet() {
	words := []string{"dog", "cat"}
	r := strings.NewReader(strings.Join(words, " "))

	dict := NewDictionary(3)
	_, _ = dict.ReadFrom(r)

	for _, r := range dict.Alphabet() {
		fmt.Println(string(r))
	}
	// Unordered output:
	// a
	// c
	// d
	// g
	// o
	// t
}

func TestDictionary_ReadFrom(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		s := "apple\nbanana  orange"
		r := strings.NewReader(s)
		dict := NewDictionary(6)
		b, err := dict.ReadFrom(r)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if b != int64(len(s)) {
			t.Errorf("number of read bytes wrong: got %v, want %v", b, len(s))
		}

		want := map[string]struct{}{
			"banana": {},
			"orange": {},
		}
		if !reflect.DeepEqual(dict.words, want) {
			t.Errorf("dict.words: got %v, want %v", dict.words, want)
		}
	})

	t.Run("ReadError", func(t *testing.T) {
		r := errReader{err: errors.New("test error")}
		dict := NewDictionary(6)
		_, err := dict.ReadFrom(r)
		if !errors.Is(err, r.err) {
			t.Errorf("dict.ReadFrom expected to return an error")
		}
		if len(dict.words) > 0 {
			t.Errorf("expected words to be empty, got %v items", len(dict.words))
		}
	})
}

func TestDictionary_AddWord(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		dict := NewDictionary(6)
		dict.AddWord("banana")

		wantWords := map[string]struct{}{"banana": {}}
		wantAlphabet := map[rune]struct{}{
			'b': {},
			'a': {},
			'n': {},
		}

		if !reflect.DeepEqual(dict.words, wantWords) {
			t.Errorf("dict.words: got %v, want %v", dict.words, wantWords)
		}
		if !reflect.DeepEqual(dict.alphabet, wantAlphabet) {
			t.Errorf("dict.alphabet: got %v, want %v", dict.alphabet, wantAlphabet)
		}
	})

	t.Run("Uppercase", func(t *testing.T) {
		dict := NewDictionary(6)
		dict.AddWord("BANANA")

		wantWords := map[string]struct{}{"banana": {}}
		wantAlphabet := map[rune]struct{}{
			'b': {},
			'a': {},
			'n': {},
		}

		if !reflect.DeepEqual(dict.words, wantWords) {
			t.Errorf("dict.words: got %v, want %v", dict.words, wantWords)
		}
		if !reflect.DeepEqual(dict.alphabet, wantAlphabet) {
			t.Errorf("dict.alphabet: got %v, want %v", dict.alphabet, wantAlphabet)
		}
	})

	t.Run("WrongLength", func(t *testing.T) {
		dict := NewDictionary(6)
		dict.AddWord("apple")

		if len(dict.words) > 0 {
			t.Errorf("expected words to be empty, got %v items", len(dict.words))
		}
		if len(dict.alphabet) > 0 {
			t.Errorf("expected alphabet to be empty, got %v items", len(dict.alphabet))
		}
	})
}

func TestDictionary_HasWord(t *testing.T) {
	dict := &Dictionary{words: map[string]struct{}{"banana": {}}}
	if !dict.HasWord("banana") {
		t.Errorf(`expected dict.HasWord to return true for "banana"`)
	}
	if dict.HasWord("apple") {
		t.Errorf(`expected dict.HasWord to return false for "apple"`)
	}
}

func TestDictionary_Alphabet(t *testing.T) {
	dict := &Dictionary{alphabet: map[rune]struct{}{
		'a': {},
		'b': {},
		'n': {},
	}}
	want := []rune{'a', 'b', 'n'}
	got := dict.Alphabet()
	// dict.Alphabet() returns an unsorted slice
	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dict.Alphabet(): got %v, want %v", got, want)
	}
}

// errReader is a helper io.Reader that always returns error err.
type errReader struct {
	err error
}
func (r errReader) Read(_ []byte) (int, error) { return 0, r.err }

