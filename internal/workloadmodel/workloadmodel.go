package workloadmodel

type Workload struct {
	Name string
	QPS  uint64
}

func NewWorkload(
	name string,
	qps uint64,
) *Workload {

	return &Workload{
		Name: name,
		QPS:  qps,
	}
}
