package main

import (
	"flag"
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	flag.Parse()

	crx, cmd := flag.Arg(0), flag.Arg(1)
	if crx == "" {
		panic("provide cron expression please")
	}
	if cmd == "" {
		panic("provide command to run please")
	}

	c := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
			),
		),
	)

	_, err := c.AddFunc(crx, func() {
		fmt.Println(cmd)
	})
	if err != nil {
		panic(err)
	}

	c.Run()
}
