package vocabulary_test

import (
	"github.com/mk2/yon/interp/vocabulary"
	"testing"
)

func TestExtractClass(t *testing.T) {

	class, key := vocabulary.ExtractClass("prelude~dup")

	if class != "prelude" || key != "dup" {
		t.Errorf("invalid extracting class: %s - %s", class, key)
	}

	class, key = vocabulary.ExtractClass("test~test~key")

	if class != "test~test" || key != "key" {
		t.Errorf("invalid extracting class: %s - %s", class, key)
	}

	class, key = vocabulary.ExtractClass("key2")

	if class != "" || key != "key2" {
		t.Errorf("invalid extracting class: %s - %s", class, key)
	}
}
