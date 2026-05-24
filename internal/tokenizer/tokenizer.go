package tokenizer

import (
	"strings"
)

type Tokenizer struct{}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Tokenize(
	input string,
) []string {

	return strings.Fields(input)
}
