package queue

type State struct {
	Coord    Coord
	Distance int
}

type Coord struct {
	Row int
	Col int
}

type Queue struct {
	items []State
}

func (q *Queue) Push(val State) {
	q.items = append(q.items, val)

}

func (q *Queue) Pop() State {
	var temp State = q.items[0]
	q.items = q.items[1:]
	return temp
}

func (q *Queue) Peek() State {
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
