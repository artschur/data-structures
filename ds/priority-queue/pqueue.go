package priorityqueue

type element struct {
	priority int
	value    any
}

type pqueue struct {
	queue []element
}

func (p *pqueue) insert(newElement *element) {
	p.queue = append(p.queue, *newElement)

	i := len(p.queue) - 1
	for i > 0 && p.queue[i-1].priority < p.queue[i].priority {
		p.queue[i], p.queue[i-1] = p.queue[i-1], p.queue[i]
		i--
	}
}
