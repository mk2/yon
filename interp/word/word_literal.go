package word

type NumberWord struct {
	BaseWord
	Number float64
}

func (w *NumberWord) Exec() (result Result, err error) {

	if _, err = w.CanExec(); err != nil {
		return
	}

	return
}

type StringWord struct {
	BaseWord
	String string
}
