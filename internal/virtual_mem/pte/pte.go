package pte

const (
	Present     uint64 = 1 << 0
	Writable    uint64 = 1 << 1
	Dirty       uint64 = 1 << 2
	Accessed    uint64 = 1 << 3
	CopyOnWrite uint64 = 1 << 4
	HugePage    uint64 = 1 << 5
	Speculative uint64 = 1 << 6
)

type Entry struct {
	Physical uint64
	Flags    uint64
}

func (e *Entry) Set(
	flag uint64,
) {

	e.Flags |= flag
}

func (e *Entry) Clear(
	flag uint64,
) {

	e.Flags &^= flag
}

func (e *Entry) Has(
	flag uint64,
) bool {

	return e.Flags&flag != 0
}
