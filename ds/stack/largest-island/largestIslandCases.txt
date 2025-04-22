package stack

type Coord struct {
	X int
	Y int
}

type Stack struct {
	items []Coord
	Len   int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]Coord, 0),
		Len:   0,
	}
}

func (s *Stack) Push(coord Coord) {
	s.items = append(s.items, coord)
	s.Len++
}

func (s *Stack) Pop() Coord {
	if s.Len == 0 {
		return Coord{}
	}
	last := s.items[s.Len-1]
	s.items = s.items[:s.Len-1]
	s.Len--
	return last
}

func (s *Stack) Peek() Coord {
	if s.Len == 0 {
		return Coord{}
	}
	return s.items[s.Len-1]
}
