package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/google/subcommands"
)

type clientCmd struct {
	protocol string
}

func (*clientCmd) Name() string     { return "client" }
func (*clientCmd) Synopsis() string { return "client args to stdout." }
func (*clientCmd) Usage() string {
	return `client:
	Run client.
  `
}

func (p *clientCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.protocol, "protocol", "http", "http or tcp")
}

func (c *clientCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.protocol == "tcp" {
		c.requestTCP()
	} else {
		c.requestHTTP()
	}
	return subcommands.ExitSuccess
}

func (p *clientCmd) requestHTTP() {
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
}

func (p *clientCmd) requestTCP() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte("hello"))

	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))
}
