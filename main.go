package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func findFnm() string {
	// Try local directory first
	execPath, err := os.Executable()
	if err == nil {
		localFnm := filepath.Join(filepath.Dir(execPath), "fnm")
		if runtime.GOOS == "windows" {
			localFnm += ".exe"
		}
		if _, err := os.Stat(localFnm); err == nil {
			fmt.Fprintf(os.Stderr, "Using local fnm: %s\n", localFnm)
			return localFnm
		}
	}

	// If we have fnm in PATH, get its actual location
	if pathFnm, err := exec.LookPath("fnm"); err == nil {
		fmt.Fprintf(os.Stderr, "Using system fnm: %s\n", pathFnm)
		return pathFnm
	}

	fmt.Fprintln(os.Stderr, "Warning: fnm not found in local directory or PATH")
	return "fnm"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: fnmx <node-version> [command]")
		os.Exit(1)
	}

	fnmPath := findFnm()
	version := os.Args[1]

	// Install the version
	install := exec.Command(fnmPath, "install", version)
	install.Stdout = os.Stderr
	install.Stderr = os.Stderr
	if err := install.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing Node version: %v\n", err)
		os.Exit(1)
	}

	// If there's a command to run
	if len(os.Args) > 2 {
		exec := exec.Command(fnmPath, append([]string{"exec", "--using", version}, os.Args[2:]...)...)
		exec.Stdout = os.Stdout
		exec.Stderr = os.Stderr
		exec.Stdin = os.Stdin
		if err := exec.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
			os.Exit(1)
		}
	}
}
