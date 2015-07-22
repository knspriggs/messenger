package messenger

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

type (
	Worker struct {
		Cmd              string //as a single string
		Cmdargs          []string
		Rundir           string
		ChanBufferLength int
	}
)

func New(cmd string, cmdargs []string, rundir string, chanBufferLength int) *Worker {
	return &Worker{
		Cmd:              cmd,
		Cmdargs:          cmdargs,
		Rundir:           rundir,
		ChanBufferLength: chanBufferLength,
	}
}

func (worker *Worker) Run() (chan string, error) {

	output_channel := make(chan string, worker.ChanBufferLength)

	if worker.Rundir != "" {
		os.Chdir(worker.Rundir)
	}
	cmd := exec.Command(worker.Cmd, worker.Cmdargs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error creating StdoutPipe for Cmd", err)
		return nil, err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			output_channel <- scanner.Text()
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Println("Error starting Cmd", err)
		return nil, err
	}

	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Println("Error waiting for command to finish", err)
		} else {
			close(output_channel)
		}
	}()

	return output_channel, nil
}
