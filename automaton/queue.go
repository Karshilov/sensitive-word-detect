package automaton

type Queue interface {
	Push(e ...int)
	Pop() int
	Front() int
	Empty() bool
	Size() int
	Clear()
}

type ArrayQueue []int

func (s *ArrayQueue) Push(e ...int) {
	*s = append(*s, e...)
}

func (s *ArrayQueue) Pop() int {
	if s.Empty() {
		return -1
	}
	first := (*s)[0]
	*s = (*s)[1:]
	return first
}

func (s *ArrayQueue) Front() int {
	if s.Empty() {
		return -1
	}
	return (*s)[0]
}

func (s *ArrayQueue) Empty() bool {
	return s.Size() == 0
}

func (s *ArrayQueue) Size() int {
	return len(*s)
}

func (s *ArrayQueue) Clear() {
	*s = ArrayQueue{}
}

func NewArrayQueue() Queue {
	return &ArrayQueue{}
}
