package raftlog

type Entry struct {
	Term uint64
	Data []byte
}

type RaftLog struct {
	entries []Entry
}

func NewRaftLog() *RaftLog {

	return &RaftLog{
		entries: make(
			[]Entry,
			0,
		),
	}
}

func (r *RaftLog) Append(
	term uint64,
	data []byte,
) {

	r.entries = append(
		r.entries,
		Entry{
			Term: term,
			Data: data,
		},
	)
}

func (r *RaftLog) Size() int {
	return len(r.entries)
}
