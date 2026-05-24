package futures

type Future struct {
	result chan int
}

func NewFuture() *Future {

	return &Future{
		result: make(chan int, 1),
	}
}

func (f *Future) Complete(
	v int,
) {

	f.result <- v
}

func (f *Future) Await() int {
	return <-f.result
}
