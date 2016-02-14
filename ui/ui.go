package ui

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Ui interface {
	Say(string)
}

type BasicUi struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
	l           sync.Mutex
	interrupted bool
}

func (rw *BasicUi) Say(message string) {
	rw.l.Lock()
	defer rw.l.Unlock()

	log.Printf("ui: %s", message)
	_, err := fmt.Fprint(rw.Writer, message+"\n")
	if err != nil {
		log.Printf("[ERR] Failed to write to UI: %s", err)
	}
}

type MachineReadableUi struct {
	Writer io.Writer
}

func (u *MachineReadableUi) Say(message string) {
	u.Machine("ui", "say", message)
}

func (u *MachineReadableUi) Machine(category string, args ...string) {
	now := time.Now().UTC()

	// Determine if we have a target, and set it
	target := ""
	commaIdx := strings.Index(category, ",")
	if commaIdx > -1 {
		target = category[0:commaIdx]
		category = category[commaIdx+1:]
	}

	// Prepare the args
	for i, v := range args {
		args[i] = strings.Replace(v, ",", "%!(PACKER_COMMA)", -1)
		args[i] = strings.Replace(args[i], "\r", "\\r", -1)
		args[i] = strings.Replace(args[i], "\n", "\\n", -1)
	}
	argsString := strings.Join(args, ",")

	_, err := fmt.Fprintf(u.Writer, "%d,%s,%s,%s\n", now.Unix(), target, category, argsString)
	if err != nil {
		if err == syscall.EPIPE || strings.Contains(err.Error(), "broken pipe") {
			// Ignore epipe errors because that just means that the file
			// is probably closed or going to /dev/null or something.
		} else {
			panic(err)
		}
	}
}
