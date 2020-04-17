package wordchain

import (
	"container/list"
)

type node struct {
	word     string
	similar  []string
	distance int
}

// nodeQueue is a simple wrapper on top of Go's container/list
type nodeQueue struct {
	l *list.List
}

func newQueue() *nodeQueue           { return &nodeQueue{l: list.New()} }
func (q *nodeQueue) len() int        { return q.l.Len() }
func (q *nodeQueue) reset()          { q.l.Init() }
func (q *nodeQueue) enqueue(n *node) { q.l.PushBack(n) }
func (q *nodeQueue) dequeue() *node {
	e := q.l.Front()
	if e == nil {
		return nil
	}
	q.l.Remove(e)
	return e.Value.(*node)
}

