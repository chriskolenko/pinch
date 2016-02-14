package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/mitchellh/cli"
	"github.com/webcanvas/pinch/ui"

	_ "github.com/webcanvas/pinch/modules/docker"
	_ "github.com/webcanvas/pinch/modules/msbuild"
)

var version string

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	os.Exit(run())
}

func run() int {
	// If there is no explicit number of Go threads to use, then set it
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// log.Printf("[INFO] Packer version: %s %s %s", Version, VersionPrerelease, GitCommit)
	log.Printf("Packer Target OS/Arch: %s %s", runtime.GOOS, runtime.GOARCH)
	log.Printf("Built with Go Version: %s", runtime.Version())

	Ui = &ui.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stdout,
	}

	args := os.Args[1:]
	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		//HelpFunc:   excludeHelpFunc(Commands, []string{"plugin"}),
		HelpWriter: os.Stdout,
		Version:    version,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
		return 1
	}

	return exitCode
}
