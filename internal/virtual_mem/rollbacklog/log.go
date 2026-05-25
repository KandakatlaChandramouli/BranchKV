package rollbacklog

import "branchkv-core/internal/virtual_mem/delta"

type Log struct {
	entries []delta.Delta
}

func NewLog() *Log {

	return &Log{}
}

func (l *Log) Append(
	d delta.Delta,
) {

	l.entries = append(
		l.entries,
		d,
	)
}

func (l *Log) Entries() []delta.Delta {

	return l.entries
}
