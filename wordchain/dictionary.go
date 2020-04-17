package wordchain

import (
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

// Dictionary holds unique words of a specified length and maintains an alphabet of characters encountered in any of
// those words.
type Dictionary struct {
	targetLen int
	words     map[string]struct{}
	alphabet  map[rune]struct{}
}

// NewDictionary returns a new Dictionary for words of a given length.
func NewDictionary(targetLen int) *Dictionary {
	return &Dictionary{
		targetLen: targetLen,
		words:     make(map[string]struct{}),
		alphabet:  make(map[rune]struct{}),
	}
}

// ReadFrom reads data from r until EOF, splitting the words using a bufio.Scanner and adding them into the Dictionary.
// The return value is the number of bytes read. Any error encountered by the scanner is also returned.
func (d *Dictionary) ReadFrom(r io.Reader) (int64, error) {
	var total int
	sc := bufio.NewScanner(r)
	sc.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		advance, token, err := bufio.ScanWords(data, atEOF)
		total += advance
		return advance, token, err
	})

	for sc.Scan() {
		d.AddWord(sc.Text())
	}
	return int64(total), sc.Err()
}

// AddWord inserts a new word into the Dictionary and expands the alphabet with characters from that word.
func (d *Dictionary) AddWord(s string) {
	if utf8.RuneCountInString(s) != d.targetLen {
		return
	}
	word := strings.ToLower(s)
	d.words[word] = struct{}{}
	for _, letter := range word {
		d.alphabet[letter] = struct{}{}
	}
}

// HasWord returns true if the given word can be found in the Dictionary.
func (d *Dictionary) HasWord(word string) bool {
	_, ok := d.words[word]
	return ok
}

// Alphabet returns a slice of runes used in all words included in the Dictionary. The characters are unique
// but not sorted.
func (d *Dictionary) Alphabet() []rune {
	letters := make([]rune, 0, len(d.alphabet))
	for r := range d.alphabet {
		letters = append(letters, r)
	}
	return letters
}
