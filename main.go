package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/robfig/cron/v3"
)

func main() {
	flag.Parse()

	crx, cmd, argv := flag.Arg(0), flag.Arg(1), flag.Args()[2:]
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
		cmd := exec.Command(cmd, argv...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}

	c.Run()
}
