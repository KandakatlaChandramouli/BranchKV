package virtual_mem

func (m *MMU) Write(logicalID uint64, index int, value float32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	desc, ok := m.virtualMap[logicalID]
	if !ok {
		return nil
	}

	page := desc.Page

	if page.Count() > 1 {
		newPage := NewPhysicalPage(
			m.nextPageID.Add(1),
			len(page.Data),
		)

		copy(newPage.Data, page.Data)

		page.DecRef()

		desc.Page = newPage
		page = newPage
	}

	page.Data[index] = value

	return nil
}
