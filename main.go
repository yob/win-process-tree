package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/sys/windows"
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
			{
				Name:    "orphans",
				Aliases: []string{"a"},
				Usage:   "create a 3-level process tree and then orphan the leaf processes",
				Action: func(c *cli.Context) error {
					cmdOrphans(c.Args().Get(0));
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


func startTree(subcommand string, nesting int) *exec.Cmd {
	cmd := exec.Command(os.Args[0], subcommand, strconv.Itoa(nesting - 1))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
		cmd1 := startTree("3tree", nesting)
		cmd2 := startTree("3tree", nesting)
		cmd3 := startTree("3tree", nesting)

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

func terminatePid(pid uint32) error {
	proc, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, pid)
	if err != nil {
		return err
	}
	err = windows.TerminateProcess(proc, 0)
	windows.CloseHandle(proc)
	return err
}

func cmdOrphans(argNesting string) error {
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
	log.Printf("orphan nesting: %d\n", nesting)

	if nesting >= 1 {
		cmd1 := startTree("orphans", nesting)
		cmd2 := startTree("orphans", nesting)
		cmd3 := startTree("orphans", nesting)

		if err := cmd1.Wait(); err != nil {
			log.Println("cmd1 wait ERROR: %v", err)
		}
		if err := cmd2.Wait(); err != nil {
			log.Println("cmd2 wait ERROR: %v", err)
		}
		if err := cmd3.Wait(); err != nil {
			log.Println("cmd3 wait ERROR: %v", err)
		}
		time.Sleep(10 * time.Second)
	} else if nesting == 0 {
		ppid := os.Getppid()
		log.Printf("leaf process, ppid: %d\n", ppid)
		err := terminatePid(uint32(ppid))
		if err != nil {
			log.Printf("Error killing ppid %d, %v", ppid, err)
		}
		time.Sleep(300 * time.Second)
	}
	return nil
}
