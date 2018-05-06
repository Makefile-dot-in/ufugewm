package main

import (
	"os/exec"
	"os"
	"errors"
	"bufio"
	"log"
	"strings"

	"github.com/queria/golang-go-xdg"
)

const AUTOSTART = "ufugewm/autostart"

// Returns an exec.Cmd used to launch $XDG_CONFIG_HOME/ufugewm/autostart.
// Returns an error if the file couldn't be read
func ExecAutostart() (pipe *bufio.Reader, err error) {
	autostartFileName, err := xdg.Config.Find(AUTOSTART)
	if err != nil {
		return
	}
	cmd := exec.Command(autostartFileName)
	err := cmd.Start()
	if err != nil {
		return
	}
	return cmd, nil
}

// Listens for commands on pipe, calls ParseCommand() on each line, and
// puts the generated events in evchan.
func ListenForCommands(pipe *bufio.Reader, evchan chan int) {
	for {
		line, err := pipe.ReadString("\n")
		if err != nil {
			log.Println(err)
		}
		event, err := ParseCommand(line)
		if err == nil {
			evchan <- 
	}
}
