package replayengine

type Record struct {
	LSN  uint64
	Data []byte
}

type ReplayEngine struct {
	log []Record
}

func NewReplayEngine() *ReplayEngine {

	return &ReplayEngine{
		log: make(
			[]Record,
			0,
		),
	}
}

func (r *ReplayEngine) Append(
	lsn uint64,
	data []byte,
) {

	r.log = append(
		r.log,
		Record{
			LSN:  lsn,
			Data: data,
		},
	)
}

func (r *ReplayEngine) Replay() int {

	return len(r.log)
}
