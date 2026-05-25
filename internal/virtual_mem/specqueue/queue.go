package specqueue

type Queue struct {
	entries []uint64
}

func NewQueue() *Queue {

	return &Queue{}
}

func (q *Queue) Push(
	id uint64,
) {

	q.entries = append(
		q.entries,
		id,
	)
}

func (q *Queue) Pop() (
	uint64,
	bool,
) {

	if len(q.entries) == 0 {
		return 0, false
	}

	value := q.entries[0]

	q.entries = q.entries[1:]

	return value, true
}
