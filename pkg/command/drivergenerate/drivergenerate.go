package drivergenerate

import (
	"github.com/mitchellh/cli"
	"flag"
	"github.com/hashicorp/consul/command/flags"
	"github.com/bestreyer/carfinder/pkg/location"
	"github.com/labstack/gommon/log"
	"context"
	"time"
	"fmt"
	"math/rand"
)

const (
	synopsis = "Database initialization"
	help     = `Usage: driver generate [option]`
)

type cmd struct {
	UI     cli.Ui
	lr     location.Repository
	amount int
	random bool
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
	c.flags.IntVar(&c.amount, "amount", 0, "Amount of drivers locations, which will be generated")
	c.flags.BoolVar(&c.random, "random", false, "Are generated drivers locations randomly or (0,0) ?")
	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); nil != err {
		return 1
	}

	results := make(chan error, c.amount)
	jobs := make(chan *location.Location, c.amount)

	c.UI.Output(fmt.Sprintf("Start generating %d drivers....", c.amount))
	for i := 0; i < 10; i++ {
		go c.createDriversLocationsWorker(jobs, results)
	}

	var generateLocation func() (*location.Location)

	if true == c.random {
		generateLocation = generateRandomLocation
	} else {
		generateLocation = generateInitLocation
	}

	for i := 0; i < c.amount; i++ {
		jobs <- generateLocation()
	}
	close(jobs)

	for i := 0; i < c.amount; i++ {
		err := <-results
		if nil != err {
			log.Fatal(err)
		}
	}

	c.UI.Output("Completed")

	return 0
}

func (c *cmd) createDriversLocationsWorker(jobs <-chan *location.Location, results chan<- error) {
	for job := range jobs {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := c.lr.Create(ctx, job)
		results <- err
	}
}

// Random Republic of Singapore latitude
func randLatitude() float64 {
	return 1.239600 + rand.Float64()*(1.478400 - 1.239600)
}

// Random Republic of Singapore longitude
func randLongitude() float64 {
	return 103.587348 + rand.Float64()*(103.978413-103.594000)
}

func generateRandomLocation() (*location.Location) {
	return &location.Location{
		Longitude: randLongitude(),
		Latitude:  randLatitude(),
		UpdatedAt: time.Now(),
	}
}

func generateInitLocation() (*location.Location) {
	return &location.Location{
		Longitude: randLongitude(),
		Latitude:  randLatitude(),
		UpdatedAt: time.Unix(0, 0),
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return c.help
}
