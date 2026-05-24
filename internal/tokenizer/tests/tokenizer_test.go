package tests

import (
	"branchkv-core/internal/tokenizer"
	"testing"
)

func TestTokenizer(
	t *testing.T,
) {

	tok := tokenizer.NewTokenizer()

	out := tok.Tokenize(
		"branch kv runtime",
	)

	if len(out) != 3 {
		t.Fatal("tokenizer failed")
	}
}
