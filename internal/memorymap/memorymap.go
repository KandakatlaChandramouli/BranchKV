package memorymap

type Region struct {
	Start uint64
	End   uint64
}

type MemoryMap struct {
	regions []Region
}

func NewMemoryMap() *MemoryMap {

	return &MemoryMap{
		regions: make([]Region, 0),
	}
}

func (m *MemoryMap) Map(
	start uint64,
	end uint64,
) {

	m.regions = append(
		m.regions,
		Region{
			Start: start,
			End:   end,
		},
	)
}

func (m *MemoryMap) Size() int {
	return len(m.regions)
}
