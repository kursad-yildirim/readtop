package main

import (
	"bufio"
	"os"
	"time"

	"1k.local/tuff/readtop"
	"1k.local/tuff/tools"
)

func main() {

	// Parse and verify input arguments
	if !readtop.ParseArgs() {
		return
	}

	// Open top file and define line scanner
	topFile, err := os.Open(readtop.InputFileName)
	tools.Check(err)
	defer topFile.Close()
	topLineScanner := bufio.NewScanner(topFile)
	err = topLineScanner.Err()
	tools.Check(err)

	// Define "top" scanner, Scan for individual top outputs,  and  display continously
	topLineScanner.Scan() // initialize scanner
	for readtop.Top.NextTop(topLineScanner) {
		tools.ClearConsole()
		readtop.Top.PrintTop(readtop.MaxLines)
		time.Sleep(readtop.Top.Freeze[readtop.Top.HasFocus])
	}
}
