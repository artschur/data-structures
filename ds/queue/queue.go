package queue

type state struct {
	Coord    coord
	Distance int
}

type coord struct {
	Row int
	Col int
}

type Queue struct {
	items []state
}

func (q *Queue) Push(val state) {
	q.items = append(q.items, val)
}

func (q *Queue) Pop() state {
	if len(q.items) == 0 {
		return state{}
	}
	var temp state = q.items[0]
	q.items = q.items[1:]
	return temp
}

func (q *Queue) Peek() state {
	if len(q.items) == 0 {
		return state{}
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
