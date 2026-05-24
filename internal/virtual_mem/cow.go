package virtual_mem

func (m *MMU) Write(
	logicalID uint64,
	index int,
	value float32,
) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	desc, ok := m.virtualMap[logicalID]

	if !ok {
		return nil
	}

	page := desc.Page

	if page.Count() > 1 {

		newData := m.arena.Allocate(len(page.Data))

		copy(newData, page.Data)

		newPage := &PhysicalPage{
			ID:   m.nextPageID.Add(1),
			Data: newData[:len(page.Data)],
		}

		newPage.RefCount.Store(1)

		page.DecRef()

		desc.Page = newPage

		page = newPage
	}

	page.Data[index] = value

	return nil
}
