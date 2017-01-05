package main

import (
	ps "github.com/mitchellh/go-ps"
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"
)

func main () {
	var (
		processes Processes
		procs	Processes
		err error
	)

	for {
		processes, err = ps.Processes()
		if err != nil {
			log.Fatal("Can't retrieve processes list: ", err)
		}
		procs = processes.FindProcessByName("parentalcontrols")
		//procs.Print()

		if len(procs) == 1 {
			syscall.Kill(procs[0].Pid(), syscall.SIGKILL)
		}

		time.Sleep(20*time.Second)
	}
}

type Processes []ps.Process
type Process struct {
	ps.Process
}

func (p *Processes) FindProcessByName (searchString string) Processes {
	var procs Processes
	for _, proc := range *p {
		if strings.Contains(strings.ToLower(proc.Executable()), strings.ToLower(searchString)) {
			procs = append(procs, proc)
		}
	}
	return procs
}

func (p *Processes) Print () {
	fmt.Println("PID\t\tExecutable name")
	fmt.Println("===\t\t===============")
	for _, proc := range *p {
		fmt.Println(proc.Pid(), "\t\t", proc.Executable())
	}
}

func (p *Process) Kill() () {
	syscall.Kill(p.Process.Pid(), syscall.SIGKILL)
}
