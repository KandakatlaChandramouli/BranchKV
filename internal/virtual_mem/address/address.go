package address

const (
	L1Shift = 22
	L2Shift = 12

	L1Mask = 0x3FF
	L2Mask = 0x3FF
)

func Split(
	virtual uint64,
) (
	uint64,
	uint64,
	uint64,
) {

	l1 := (virtual >> L1Shift) & L1Mask

	l2 := (virtual >> L2Shift) & L2Mask

	offset := virtual & 0xFFF

	return l1, l2, offset
}
