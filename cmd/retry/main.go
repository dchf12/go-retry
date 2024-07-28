package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dchf12/go-retry"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "retry command",
		Usage: "A command line tool to retry a shell command",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "maxRetries",
				Value: 3,
				Usage: "Maximum number of retries",
			},
			&cli.DurationFlag{
				Name:  "delay",
				Value: 5 * time.Second,
				Usage: "Delay between retries (e.g., 1s, 500ms)",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				cli.ShowAppHelpAndExit(c, 1)
			}

			command := c.Args().Get(0)
			args := c.Args().Slice()[1:]

			maxRetries := c.Int("maxRetries")
			delay := c.Duration("delay")

			myFunc := func() error {
				cmd := exec.Command(command, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			}

			err := retry.Retry(myFunc,
				retry.WithMaxRetries(maxRetries),
				retry.WithDelay(delay),
			)
			if err != nil {
				fmt.Println("All retries failed:", err)
				return cli.Exit("Failed to complete the task", 1)
			}

			fmt.Println("Success!")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
