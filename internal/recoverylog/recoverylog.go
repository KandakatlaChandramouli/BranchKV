package recoverylog

type Record struct {
	LSN  uint64
	Data []byte
}

type Log struct {
	records []Record
}

func NewLog() *Log {

	return &Log{
		records: make(
			[]Record,
			0,
		),
	}
}

func (l *Log) Append(
	lsn uint64,
	data []byte,
) {

	l.records = append(
		l.records,
		Record{
			LSN:  lsn,
			Data: data,
		},
	)
}

func (l *Log) Size() int {
	return len(l.records)
}
