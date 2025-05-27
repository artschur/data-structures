package stack

type coord struct {
	X int
	Y int
}

type stack struct {
	items []coord
	Len   int
}

func NewStack() *stack {
	return &stack{
		items: make([]coord, 0),
		Len:   0,
	}
}

func (s *stack) Push(coord coord) {
	s.items = append(s.items, coord)
	s.Len++
}

func (s *stack) Pop() coord {
	if s.Len == 0 {
		return coord{}
	}
	last := s.items[s.Len-1]
	s.items = s.items[:s.Len-1]
	s.Len--
	return last
}

func (s *stack) Peek() coord {
	if s.Len == 0 {
		return coord{}
	}
	return s.items[s.Len-1]
}

func (s *stack) IsEmpty() bool {
	return s.Len == 0
}
