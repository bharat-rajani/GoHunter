package main

import (
	"github.com/bharat-rajani/GoHunter/pkg/cmd/factory"
	"github.com/bharat-rajani/GoHunter/pkg/cmd/root"
	"os"
)

type exitCode int

const (
	exitOk     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
)

func main() {
	runExitCode := runMain()
	os.Exit(int(runExitCode))

}

func runMain() exitCode {
	cmdContainerFactory := factory.New()
	rootCmd := root.NewCmdRoot(cmdContainerFactory)
	err := rootCmd.Execute()
	if err != nil {
		return exitError
	}

	return exitOk
}
