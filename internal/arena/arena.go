package arena

type Arena struct {
	Small  *Slab
	Medium *Slab
	Large  *Slab
}

func NewArena() *Arena {

	return &Arena{
		Small:  NewSlab(256, 1024),
		Medium: NewSlab(1024, 512),
		Large:  NewSlab(4096, 128),
	}
}

func (a *Arena) Allocate(size int) []float32 {

	switch {

	case size <= 256:
		return a.Small.Allocate()

	case size <= 1024:
		return a.Medium.Allocate()

	default:
		return a.Large.Allocate()
	}
}

func (a *Arena) Free(page []float32) {

	size := len(page)

	switch {

	case size <= 256:
		a.Small.Free(page)

	case size <= 1024:
		a.Medium.Free(page)

	default:
		a.Large.Free(page)
	}
}
