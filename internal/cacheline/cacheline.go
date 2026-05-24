package cacheline

const Size = 64

type Aligned struct {
	_ [Size]byte
}
