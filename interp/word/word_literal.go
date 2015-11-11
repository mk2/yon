package word

import (
	"fmt"
	"strconv"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type numberWord struct {
	word
	number float64
}

func NewNumberWord(val string) kit.NumberWord {

	var (
		f   float64
		err error
	)
	if f, err = strconv.ParseFloat(val, 64); err != nil {
		f = 0.0
	}

	return &numberWord{
		word:   word{wordType: TNumberWord, author: author.NewUserAuthor()},
		number: f,
	}
}

func (w *numberWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *numberWord) Number() float64 {

	return w.number
}

func (w *numberWord) String() string {

	return strconv.FormatFloat(w.number, 'E', -1, 64)
}

func (w *numberWord) Format() string {

	return fmt.Sprintf(fNumberWord, w.number)
}

type stringWord struct {
	word
	str string
}

func NewStringWord(val string) kit.StringWord {

	return &stringWord{
		str:  val,
		word: word{wordType: TStringWord, author: author.NewUserAuthor()},
	}
}

func (w *stringWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *stringWord) String() string {

	return w.str
}

func (w *stringWord) Format() string {

	return fmt.Sprintf(fStringWord, w.str)
}
