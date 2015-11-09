package kit

import (
	"log"

	"github.com/mk2/yon/interp/env"
)

func Printf(format string, args ...interface{}) {

	dbg.printf(format, args...)
}

func Println(args ...interface{}) {

	dbg.println(args...)
}

type debugT bool

var dbg debugT = debugT(env.Debug)

func (d debugT) printf(format string, args ...interface{}) {

	if d {
		log.Printf(format, args...)
	}
}

func (d debugT) println(args ...interface{}) {

	if d {
		log.Println(args...)
	}
}
