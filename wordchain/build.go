package wordchain

import (
	"errors"
)

// Errors that can be returned from the Build function.
var (
	// ErrEmpty is returned when both of the words are empty strings.
	ErrEmpty = errors.New("both words must not be empty")

	// ErrNotInDictionary is returned when either of the words could not be found in the provided dictionary.
	ErrNotInDictionary = errors.New("both words must be present in the dictionary")

	// ErrNotEqualLength is returned when the word lengths differ. Since a Dictionary can only contain words of
	// a specific length, both words must be of this length as well.
	ErrNotEqualLength = errors.New("both words must be the same length")

	// ErrNotFound is returned if a chain cannot be built as there is no possible path from origin to target word using
	// words in the dictionary.
	ErrNotFound = errors.New("chain not found")
)

// Build finds the shortest chain of words from origin to target by changing only one character in a word each time.
// All words in the chain must be included in the passed dictionary.
func Build(dict *Dictionary, origin, target string) ([]string, error) {
	if len(origin) != len(target) {
		return nil, ErrNotEqualLength
	}
	if len(origin) == 0 {
		return nil, ErrEmpty
	}
	if !dict.HasWord(origin) || !dict.HasWord(target) {
		return nil, ErrNotInDictionary
	}

	b := newBuilder(dict, origin, target)
	b.populateNodes()
	return b.shortestChain()
}

// builder is a helper struct used to find the word chain, holding the dictionary, origin and target words, and
// a map of visited nodes
type builder struct {
	dict            *Dictionary
	origin, target string
	nodes           map[string]*node
}

func newBuilder(dict *Dictionary, origin, target string) *builder {
	b := &builder{
		dict:    dict,
		origin: origin,
		target:  target,
		nodes:   make(map[string]*node),
	}
	b.nodes[origin] = &node{word: origin}
	return b
}

// populateNodes uses breadth-first search to find the target word, adding all the visited nodes to the nodes map.
func (b *builder) populateNodes() {
	q := newQueue()
	q.enqueue(b.nodes[b.origin])

	for q.len() > 0 {
		current := q.dequeue()
		if current == nil {
			break
		}
		current.similar = b.similarWords(current.word)
		for _, word := range current.similar {
			if _, ok := b.nodes[word]; !ok {
				b.nodes[word] = &node{word: word, distance: current.distance + 1}
				if word == b.target {
					// store the target node's similar words here, as after resetting the queue the outer loop
					// will break
					b.nodes[word].similar = b.similarWords(word)
					q.reset()
				} else {
					q.enqueue(b.nodes[word])
				}
			}
		}
	}
}

// similarWords finds words from the dictionary that are similar (e.g. with only one different character) to the passed
// word. It does this by exchanging each character with all possible characters from the alphabet and saving the
// combinations that are present in the dictionary.
func (b *builder) similarWords(word string) []string {
	similarMap := make(map[string]bool)
	runes := []rune(word)
	for i := range runes {
		orig := runes[i]
		for _, letter := range b.dict.Alphabet() {
			runes[i] = letter
			s := string(runes)
			if b.dict.HasWord(s) {
				similarMap[s] = true
			}
			runes[i] = orig
		}
	}
	similar := make([]string, 0, len(similarMap))
	for w := range similarMap {
		similar = append(similar, w)
	}
	return similar
}

// shortestChain returns the shortest possible chain from the origin word to the target word by using the previously
// populated nodes map. It does this by first retrieving the target word, finding a similar word that has the lowest
// distance to origin and iterating this way until reaching the origin itself.
func (b *builder) shortestChain() ([]string, error) {
	current, ok := b.nodes[b.target]
	if !ok {
		return nil, ErrNotFound
	}
	path := make([]string, current.distance+1)
	path[current.distance] = current.word
	for current.word != b.origin {
		var closest *node
		for _, w := range current.similar {
			n := b.nodes[w]
			if n != nil && (closest == nil || n.distance < closest.distance) {
				closest = n
			}
		}
		if closest == nil {
			return nil, ErrNotFound
		}
		current = closest
		path[current.distance] = current.word
	}
	return path, nil
}
