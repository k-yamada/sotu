package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/google/subcommands"
)

type clientCmd struct {
	capitalize bool
}

func (*clientCmd) Name() string     { return "client" }
func (*clientCmd) Synopsis() string { return "client args to stdout." }
func (*clientCmd) Usage() string {
	return `client [-capitalize] <some text>:
	client args to stdout.
  `
}

func (p *clientCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *clientCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()

	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://0.0.0.0:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(
		bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	return subcommands.ExitSuccess
}
