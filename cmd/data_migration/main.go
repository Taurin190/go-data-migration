package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Taurin190/go-data-migration"
	"github.com/Taurin190/go-data-migration/service"
)

const usageMessage = "" +
	`Usage:	data_migration [flags]
	data_migration is data migration script between databases.`

type option struct {
	version   bool
	srcHost   string
	srcPort   int
	srcDB     string
	srcDriver string
}

const (
	srcHostDoc   = `source host to retrieve data and schema. `
	srcPortDoc   = `source port to retrieve data and schema. `
	srcDBDoc     = `source database to retrieve data and schema`
	srcDriverDoc = `source database driver. you can choose mysqlDB, mariaDB and mongoDB.`
)

func init() {
	flag.BoolVar(&opt.version, "version", false, "print version")
	flag.StringVar(&opt.srcHost, "src-host", "localhost", srcHostDoc)
	flag.IntVar(&opt.srcPort, "src-port", 3306, srcPortDoc)
	flag.StringVar(&opt.srcDB, "src-db", "", srcDBDoc)
	flag.StringVar(&opt.srcDriver, "src-driver", "mysql", srcDriverDoc)
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
		fmt.Fprintf(os.Stderr, "data_migration: %v\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, opt *option) error {
	ctx := context.Background()

	if opt.version {
		fmt.Fprintln(w, "0.0.1")
		return nil
	}

	var sds data_migration.DBService

	switch opt.srcDriver {
	default:
		return fmt.Errorf("unknown -src-driver: %s", opt.srcDriver)
	case "mysql":
		fmt.Fprintf(w, "selected %s driver for src database\n", opt.srcDriver)
		sds = mysqlservice.CreateMySQLService()
	case "mongo":
		fmt.Fprintf(w, "selected %s driver for src database\n", opt.srcDriver)
	case "maria":
		fmt.Fprintf(w, "selected %s driver for src database\n", opt.srcDriver)
	}

	return sds.FetchSchema(ctx)
}
