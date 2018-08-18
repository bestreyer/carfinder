package generate_drivers

import (
	"github.com/mitchellh/cli"
	"flag"
	"github.com/hashicorp/consul/command/flags"
	"github.com/bestreyer/carfinder/pkg/location"
	"github.com/labstack/gommon/log"
)

const (
	synopsis = "Database initialization"
	help     = `Usage: driver generate`
)

type cmd struct {
	UI     cli.Ui
	lr     location.Repository
	amount int
	flags  *flag.FlagSet
	help   string
}

func New(ui cli.Ui, lr location.Repository) *cmd {
	c := &cmd{UI: ui, lr: lr}
	c.init()
	return c
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.IntVar(&c.amount, "amount", 0, "Amount of drivers, which will be generated")
	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); nil != err {
		return 1
	}

	c.UI.Output("Start generating drivers....")

	for i := 0; i < c.amount; i++ {
		go c.generateDriver()
	}

	c.UI.Output("Completed")

	return 0
}

func (c *cmd) generateDriver() {
	l := location.Location{}
	err := c.lr.Create(&l)
	if nil != err {
		log.Fatal(err)
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}
