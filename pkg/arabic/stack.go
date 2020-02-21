package arabic

//Stack is used to aggregate Non-Arabic words
type Stack struct {
	stack    []string
	capacity int
	size     int
}

// NewStack creates a new Stack
func NewStack(cap int) *Stack {
	return &Stack{
		stack:    make([]string, cap),
		capacity: cap,
		size:     0,
	}
}

// Push appends an element to the stack
func (s *Stack) Push(word string) {
	if s.size >= s.capacity {
		panic("Stack overflow")
	}

	s.stack[s.size] = word
	s.size++
}

// Pop removes the top element from the stack
func (s *Stack) Pop() string {
	if s.size == 0 {
		return ""
	}

	s.size--
	return s.stack[s.size]
}

// Flush empties the stack returing all the elements
func (s *Stack) Flush() []string {
	result := make([]string, s.size)

	i := 0
	for e := s.Pop(); e != ""; e = s.Pop() {
		result[i] = e
		i++
	}

	return result
}
