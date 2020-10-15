package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Name: "win-process-tree",
		Usage: "Create various shaped process trees. Useful for testing process management code",
		Commands: []*cli.Command{
			{
				Name:    "3tree",
				Aliases: []string{"a"},
				Usage:   "create a 3-level process tree that runs for about 30 seconds",
				Action: func(c *cli.Context) error {
					cmdThreeTree(c.Args().Get(0));
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}


func startTree(nesting int) *exec.Cmd {
	cmd := exec.Command(os.Args[0], "3tree", strconv.Itoa(nesting - 1))
	if err := cmd.Start(); err != nil {
		log.Println("Start ERROR: %v", err)
	}

	return cmd
}


func cmdThreeTree(argNesting string) error {
	nesting := 0

	if argNesting == "" { // # top of the process tree
		nesting = 3
	} else {
		res, err := strconv.Atoi(argNesting)
		if err != nil {
			nesting = 1
		}
		nesting = res
	}
	log.Printf("nesting: %d\n", nesting)

	if nesting >= 1 {
		cmd1 := startTree(nesting)
		cmd2 := startTree(nesting)
		cmd3 := startTree(nesting)

		if err := cmd1.Wait(); err != nil {
			log.Println("cmd1 wait ERROR: %v", err)
		}
		if err := cmd2.Wait(); err != nil {
			log.Println("cmd2 wait ERROR: %v", err)
		}
		if err := cmd3.Wait(); err != nil {
			log.Println("cmd3 wait ERROR: %v", err)
		}
	}
	time.Sleep(10 * time.Second)
	return nil
}
