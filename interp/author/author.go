package author

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"

	"github.com/mk2/yon/interp/kit"
)

type author struct {
	typ kit.AuthorType
	id  kit.AuthorId
}

const (
	AuthorPrelude kit.AuthorType = "prelude"
	AuthorGo      kit.AuthorType = "go"
	AuthorUser    kit.AuthorType = "user"
)

func randomAuthorId() kit.AuthorId {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return kit.AuthorId(strconv.FormatUint(n, 36))
}

func NewPreludeAuthor() kit.Author {
	return NewStaticAuthor(AuthorPrelude)
}

func NewUserAuthor() kit.Author {
	return NewRandomAuthor(AuthorUser)
}

func NewStaticAuthor(typ kit.AuthorType) kit.Author {
	return New(typ, "")
}

func NewRandomAuthor(typ kit.AuthorType) kit.Author {
	return New(typ, randomAuthorId())
}

func New(typ kit.AuthorType, id kit.AuthorId) kit.Author {
	return &author{
		typ: typ,
		id:  id,
	}
}

func (a *author) GetAuthorType() kit.AuthorType {
	return a.typ
}

func (a *author) GetAuthorId() kit.AuthorId {
	return a.id
}
