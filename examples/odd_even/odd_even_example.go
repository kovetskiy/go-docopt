package main

import (
	"fmt"
	"github.com/kovetskiy/godocs"
)

func main() {
	usage := `Usage: odd_even_example [-h | --help] (ODD EVEN)...

Example, try:
  odd_even_example 1 2 3 4

Options:
  -h, --help`

	arguments, _ := godocs.Parse(usage, nil, true, "", false)
	fmt.Println(arguments)
}
