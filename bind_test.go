package godocs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	test := assert.New(t)

	value := struct {
		Key   bool  `arg:"--key"`
		Value int64 `arg:"-v"`
	}{}

	err := Bind(`
Usage:
	p [options]

Options:
  -k --key    Key
  -v <value>  Value
`, `1`, &value, Args{"-v", "1"})
	test.NoError(err)
	fmt.Printf("XXXXXX bind_test.go:26: value: %#v\n", value)

}
