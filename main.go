package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// We use this struct to retreive process handle(which is unexported)
// from os.Process using unsafe operation.
type process struct {
	Pid    int
	Handle uintptr
}

type ProcessExitGroup windows.Handle

func NewProcessExitGroup() (ProcessExitGroup, error) {
	handle, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return 0, err
	}

	info := windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION{
		BasicLimitInformation: windows.JOBOBJECT_BASIC_LIMIT_INFORMATION{
			LimitFlags: windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE,
		},
	}
	if _, err := windows.SetInformationJobObject(
		handle,
		windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info))); err != nil {
		return 0, err
	}

	return ProcessExitGroup(handle), nil
}

func (g ProcessExitGroup) Dispose() error {
	return windows.CloseHandle(windows.Handle(g))
}

func (g ProcessExitGroup) AddProcess(p *os.Process) error {
	return windows.AssignProcessToJobObject(
		windows.Handle(g),
		windows.Handle((*process)(unsafe.Pointer(p)).Handle))
}

func startTree(nesting int) *exec.Cmd {
	//g, err := NewProcessExitGroup()
	//if err != nil {
	//	panic(err)
	//}
	//defer g.Dispose()

	cmd := exec.Command(os.Args[0], strconv.Itoa(nesting - 1))
	if err := cmd.Start(); err != nil {
		log.Println("Start ERROR: %v", err)
	}

	//if err := g.AddProcess(cmd.Process); err != nil {
	//	panic(err)
	//}
	return cmd

}

func main() {
	argCount := len(os.Args)
	if argCount != 2 {
		panic("USAGE: foo <n>")
	}
	nesting, err := strconv.Atoi(os.Args[1])
	if err != nil {
		nesting = 1
	}
	log.Printf("nesting: %d\n", nesting)

	if nesting >= 1 {
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
