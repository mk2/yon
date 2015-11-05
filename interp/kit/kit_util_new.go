package kit

// NewStoppedCh returns new stopped channel
func NewStoppedCh() StoppedCh {

	return make(StoppedCh)
}

// NewErrorCh returns new error channel
func NewErrorCh() ErrorCh {

	return make(ErrorCh)
}
