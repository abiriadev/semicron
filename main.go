package main

import (
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/robfig/cron/v3"
)

type Cfg struct {
	Sh   string `short:"s" long:"shell" description:"Shell to execute the command"`
	Args struct {
		Cron string   `positional-arg-name:"cron"`
		Cmd  string   `positional-arg-name:"cmd"`
		Argv []string `positional-arg-name:"arguments"`
	} `                                                                  positional-args:"yes" required:"yes"`
}

func main() {
	var cfg Cfg
	_, err := flags.Parse(&cfg)
	if err != nil {
		os.Exit(1)
	}

	c := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
			),
		),
	)

	resolvedCmdName, err := exec.LookPath(cfg.Args.Cmd)
	if err != nil {
		panic(err)
	}

	_, err = c.AddFunc(cfg.Args.Cron, func() {
		var cmd *exec.Cmd
		if cfg.Sh != "" {
			cmd = exec.Command(
				cfg.Sh,
				"-c",
				strings.Join(append([]string{resolvedCmdName}, cfg.Args.Argv...), " "),
			)
		} else {
			cmd = exec.Command(resolvedCmdName, cfg.Args.Argv...)
		}
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		c.Stop()
	}()

	c.Run()
}
