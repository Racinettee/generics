package generics

type Queue[X any] List[X]

func (queue Queue[X]) Front() X {
	return queue[0]
}

func (queue Queue[X]) Back() X {
	return queue[len(queue)-1]
}

func (queue *Queue[X]) Pop() X {
	res := queue.Front()
	*queue = (*queue)[1:]
	return res
}

func NewQueue[X any](size int) Queue[X] {
	return make([]X, size)
}
