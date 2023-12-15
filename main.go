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

	c := cron.New()

	_, err := c.AddFunc(crx, func() {
		fmt.Println(cmd)
	})
	if err != nil {
		panic(err)
	}

	c.Run()
}
