package history

import "github.com/mk2/yon/interp/kit"

type history struct {

}

func New() kit.History {

	return &history{

	}
}

func (h *history) Leave(w kit.Word) error {

	return nil
}
