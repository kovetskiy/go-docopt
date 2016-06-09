package main

import (
	"fmt"
	"github.com/godoc/godoc"
)

func main() {
	usage := `Usage:
  quick_example tcp <host> <port> [--timeout=<seconds>]
  quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
  quick_example -h | --help | --version`

	arguments, _ := godoc.Parse(usage, nil, true, "0.1.1rc", false)
	fmt.Println(arguments)
}
