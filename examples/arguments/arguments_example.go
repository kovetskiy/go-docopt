package main

import (
	"fmt"
	"github.com/kovetskiy/godocs"
)

func main() {
	usage := `Usage: arguments_example [-vqrh] [FILE] ...
       arguments_example (--left | --right) CORRECTION FILE

Process FILE and optionally apply correction to either left-hand side or
right-hand side.

Arguments:
  FILE        optional input file
  CORRECTION  correction angle, needs FILE, --left or --right to be present

Options:
  -h --help
  -v       verbose mode
  -q       quiet mode
  -r       make report
  --left   use left-hand side
  --right  use right-hand side`

	arguments, _ := godocs.Parse(usage, nil, true, "", false)
	fmt.Println(arguments)
}
