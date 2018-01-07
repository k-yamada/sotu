package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/google/subcommands"
)

type serverCmd struct {
	protocol string
}

func (*serverCmd) Name() string     { return "server" }
func (*serverCmd) Synopsis() string { return "server args to stdout." }
func (*serverCmd) Usage() string {
	return `server:
	Run server
	`
}

func (p *serverCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.protocol, "protocol", "http", "http or tcp")
}

func (p *serverCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s Server is running at 0.0.0.0:8888\n", strings.ToUpper(p.protocol))
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		if p.protocol == "tcp" {
			go p.handleTCP(conn)
		} else {
			go p.handleHTTP(conn)
		}
	}

	return subcommands.ExitSuccess
}

func (s *serverCmd) handleHTTP(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// リクエストを読み込む
	request, err := http.ReadRequest(
		bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	// レスポンスを書き込む
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body: ioutil.NopCloser(
			strings.NewReader("HTTP connection was successful\n")),
	}
	response.Write(conn)
	conn.Close()
}

func (s *serverCmd) handleTCP(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	conn.Write([]byte("Hello"))
	conn.Close()
}
