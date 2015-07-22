package messenger

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

type (
	Messenger struct {
		Cmd              string //as a single string
		Cmdargs          []string
		Rundir           string
		ChanBufferLength int
	}
)

func New(cmd string, cmdargs []string, rundir string, chanBufferLength int) *Messenger {
	return &Messenger{
		Cmd:              cmd,
		Cmdargs:          cmdargs,
		Rundir:           rundir,
		ChanBufferLength: chanBufferLength,
	}
}

func (messenger *Messenger) Run() (chan string, error) {

	output_channel := make(chan string, messenger.ChanBufferLength)

	if messenger.Rundir != "" {
		os.Chdir(messenger.Rundir)
	}
	cmd := exec.Command(messenger.Cmd, messenger.Cmdargs...)
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
