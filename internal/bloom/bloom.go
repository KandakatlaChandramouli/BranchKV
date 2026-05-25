package bloom

type Bloom struct {
	bits map[string]bool
}

func NewBloom() *Bloom {

	return &Bloom{
		bits: make(
			map[string]bool,
		),
	}
}

func (b *Bloom) Add(
	key string,
) {

	b.bits[key] = true
}

func (b *Bloom) Contains(
	key string,
) bool {

	return b.bits[key]
}
