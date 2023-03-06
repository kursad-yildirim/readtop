package main

import (
	"bufio"
	"os"
	"time"

	"redhat.com/ess/esstools"
	"redhat.com/ess/readtop"
)

func main() {

	// Parse and verify input arguments
	if !readtop.ParseArgs() {
		return
	}

	// Open top file and define line scanner
	topFile, err := os.Open(readtop.InputFileName)
	esstools.Check(err)
	defer topFile.Close()
	topLineScanner := bufio.NewScanner(topFile)
	err = topLineScanner.Err()
	esstools.Check(err)

	// Define "top" scanner, Scan for individual top outputs,  and  display continously
	topLineScanner.Scan() // initialize scanner
	for readtop.Top.NextTop(topLineScanner) {
		esstools.ClearConsole()
		readtop.Top.PrintTop(readtop.MaxLines)
		time.Sleep(readtop.Top.Freeze[readtop.Top.HasFocus])
	}
}
