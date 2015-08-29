package word

import "github.com/mk2/yon/interp/kit"

const (
	TNilWord kit.WordType = iota
	TFuncWord
	TNumberWord
	TStringWord
	TArrayWord
)
