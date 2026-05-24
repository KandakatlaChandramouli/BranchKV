package tensorparallel

type Shard struct {
	ID uint64
}

type TensorParallel struct {
	shards []Shard
}

func NewTensorParallel() *TensorParallel {

	return &TensorParallel{
		shards: make([]Shard, 0),
	}
}

func (t *TensorParallel) AddShard(
	id uint64,
) {

	t.shards = append(
		t.shards,
		Shard{
			ID: id,
		},
	)
}

func (t *TensorParallel) Size() int {
	return len(t.shards)
}
