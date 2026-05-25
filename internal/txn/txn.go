package txn

type Transaction struct {
	ID     uint64
	Active bool
}

func Begin(
	id uint64,
) *Transaction {

	return &Transaction{
		ID:     id,
		Active: true,
	}
}

func (t *Transaction) Commit() {
	t.Active = false
}
