package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	//"time"
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


type JOBOBJECT_BASIC_PROCESS_ID_LIST struct {
	NumberOfAssignedProcesses uint32
	NumberOfProcessIdsInList  uint32
	ProcessIdList [1024]byte
}

func (g ProcessExitGroup) ListProcesses() ([]uint64, error) {
	JobObjectBasicProcessIdList := int32(3)
	var list JOBOBJECT_BASIC_PROCESS_ID_LIST


	err := windows.QueryInformationJobObject(
		windows.Handle(g),
	    JobObjectBasicProcessIdList,
		uintptr(unsafe.Pointer(&list)),
		uint32(unsafe.Sizeof(list)),
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("QueryInformationJobObject failed: %v", err))
	}

	if  list.NumberOfProcessIdsInList < list.NumberOfAssignedProcesses {
		return nil, errors.New(fmt.Sprintf("JOBOBJECT_BASIC_PROCESS_ID_LIST buffer not large enough. Got %d pids, wanted %d", list.NumberOfProcessIdsInList, list.NumberOfAssignedProcesses))
	}

	pids := make([]uint64, list.NumberOfProcessIdsInList)
	for i := 0; i < int(list.NumberOfProcessIdsInList); i++ {
		pids[i] = binary.LittleEndian.Uint64(list.ProcessIdList[8*i : 8*(i+1)])
	}
	fmt.Printf("bnums: %v\n", pids)
	return pids, nil
}

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
		g, err := NewProcessExitGroup()
		if err != nil {
			panic(err)
		}
		//defer g.Dispose()
		for i := 1; i < 100; i++ {
			cmd1 := startTree(nesting)
			if err := g.AddProcess(cmd1.Process); err != nil {
				panic(err)
			}

			if err := cmd1.Wait(); err != nil {
				log.Println("cmd1 wait ERROR: %v", err)
			}
		}

		g.ListProcesses()

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
	//time.Sleep(500 * time.Millisecond)
}
