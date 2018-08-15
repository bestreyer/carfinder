package server

import (
	"github.com/mitchellh/cli"
	"flag"
	"github.com/hashicorp/consul/command/flags"
	"fmt"
	"github.com/bestreyer/carfinder/server"
	"log"
	"os"
)

const (
	synopsis  = "Start server"
	addrUsage = "The address to listen to (can be address:port, address, or port, default: %s)"
	help      = `Usage: server start [options]
Start http server: By default, the server listens on %s address
`
)

type cmd struct {
	UI     cli.Ui
	server server.HTTPServerInterface
	addr   string
	flags  *flag.FlagSet
	help   string
}

func New(ui cli.Ui, hs server.HTTPServerInterface) *cmd {
	c := &cmd{UI: ui, server: hs}
	c.init()
	return c
}

func (c *cmd) init() {
	defaultAddr := os.Getenv("CARFINDER_ADDR")
	if defaultAddr == "" {
		defaultAddr = "127.0.0.1"
	}

	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.addr, "addr", defaultAddr, fmt.Sprintf(addrUsage, defaultAddr))
	c.help = flags.Usage(fmt.Sprintf(help, defaultAddr), c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); nil != err {
		return 1
	}

	if nil == c.server {
		log.Fatal("c.server should not be nil.")
		return 2
	}

	c.UI.Output(fmt.Sprintf("Start server on %s", c.addr))
	if err := c.server.Start(c.addr); nil != err {
		log.Fatal(err)
		return 3
	}

	return 0;
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return c.help
}
