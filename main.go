package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)


func startTree(nesting int) *exec.Cmd {
	cmd := exec.Command(os.Args[0], strconv.Itoa(nesting - 1))
	if err := cmd.Start(); err != nil {
		log.Println("Start ERROR: %v", err)
	}

	return cmd
}

func main() {
	argCount := len(os.Args)
	nesting := 0
	err := errors.New("fo")
	if argCount == 1 { // # top of the process tree
		nesting = 3
	} else {
		nesting, err = strconv.Atoi(os.Args[1])
		if err != nil {
			nesting = 1
		}
	}
	log.Printf("nesting: %d\n", nesting)

	// only create a process group at the top of the tree, in theory
	// all children will inherit it
	if nesting == 3 {
		cmd1 := startTree(nesting)
		cmd2 := startTree(nesting)

		time.Sleep(2 * time.Second)

		if err := cmd1.Wait(); err != nil {
			log.Println("cmd1 wait ERROR: %v", err)
		}
		if err := cmd2.Wait(); err != nil {
			log.Println("cmd2 wait ERROR: %v", err)
		}
	} else if nesting >= 1 {
		cmd1 := startTree(nesting)
		cmd2 := startTree(nesting)

		if err := cmd1.Wait(); err != nil {
			log.Println("cmd1 wait ERROR: %v", err)
		}
		if err := cmd2.Wait(); err != nil {
			log.Println("cmd2 wait ERROR: %v", err)
		}
	}
	time.Sleep(10 * time.Second)
}
