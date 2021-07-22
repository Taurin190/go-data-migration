package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const usageMessage = "" +
	`Usage:	data_migration [flags]
	data_migration is data migration script between databases.`

type option struct {
	version bool
}

func init() {
	flag.BoolVar(&opt.version, "version", false, "print version")
}

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
}

var opt = &option{}

func main() {
	flag.Usage = usage
	flag.Parse()
	if err := run(os.Stdin, os.Stdout, opt); err != nil {
		fmt.Fprintf(os.Stderr, "reviewdog: %v\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, opt *option) error {
	if opt.version {
		fmt.Fprintln(w, "0.0.1")
		return nil
	}
	return nil
}
