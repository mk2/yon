package kit

import "time"

const (
	// LexerTimeout defines lexer default timeout seconds
	LexerTimeout time.Duration = 30 * time.Second
	// ParserTimeout defines parser default timeout seconds
	ParserTimeout time.Duration = 30 * time.Second
)
