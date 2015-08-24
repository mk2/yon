package main
import (
	"github.com/mk2/yon"
	"fmt"
)

func main() {

	interp := yon.NewInterpreter()
	e := interp.ParseAndExec(`123 456 "tesT" test? .s`)
	fmt.Println("error: ", e)
}
