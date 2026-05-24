package runtime

import (
	"branchkv-core/internal/types"
)

type TokenLog struct {
	Tokens []types.TokenID
}

func NewTokenLog(
	size int,
) *TokenLog {

	return &TokenLog{
		Tokens: make([]types.TokenID, 0, size),
	}
}

func (t *TokenLog) Append(
	token types.TokenID,
) {

	t.Tokens = append(t.Tokens, token)
}

func (t *TokenLog) Fork() *TokenLog {

	return &TokenLog{
		Tokens: t.Tokens,
	}
}
