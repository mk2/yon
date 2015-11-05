package history

import (
	"github.com/mk2/yon/interp/kit"
	"time"
)

type history struct {
}

func New() kit.History {

	return &history{}
}

func (h *history) Record(w kit.Word) error {

	return nil
}

func (h *history) RecordAt(w kit.Word, at time.Time) error {

	return nil
}

func (h *history) Between(from time.Time, to time.Time) []kit.Word {

	return nil
}
