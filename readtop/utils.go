package readtop

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func ParseArgs() bool {
	flag.Float64Var(&normalRefreshDelay, "r", 0.1, "Normal refresh delay in seconds")
	flag.Float64Var(&normalRefreshDelay, "refresh", 0.1, "Normal refresh delay in seconds")
	flag.Float64Var(&slowRefreshDelay, "s", 1, "Slow refresh delay in seconds")
	flag.Float64Var(&slowRefreshDelay, "slow", 1, "Slow refresh delay in seconds")
	flag.StringVar(&tokenToLook, "t", "", "Process Name to focus")
	flag.StringVar(&tokenToLook, "token", "", "Process Name to focus")
	flag.IntVar(&MaxLines, "l", 50, "Number of lines to display")
	flag.IntVar(&MaxLines, "lines", 50, "Number of lines to display")
	flag.BoolVar(&version, "v", false, "Application Version")
	flag.BoolVar(&version, "version", false, "Application Version")
	flag.BoolVar(&help, "h", false, "Help message.")
	flag.BoolVar(&help, "help", false, "Help message.")
	flag.Parse()
	argResult := false

	Top.Freeze = make(map[bool]time.Duration)
	Top.Freeze[false] = time.Duration(normalRefreshDelay*1000) * time.Millisecond
	Top.Freeze[true] = time.Duration(math.Max(normalRefreshDelay, slowRefreshDelay)*1000) * time.Millisecond

	switch {
	case help:
		displayHelp()
	case version:
		displayVersion()
	case len(flag.Args()) == 1:
		InputFileName = flag.Args()[0]
		argResult = true
	default:
		fmt.Println("\nERROR:\tNo File name is provided. Filename is mandatory!")
		displayHelp()
	}
	return argResult
}

func displayHelp() {
	fmt.Println("\nread-top:\n\tDisplays linux 'top' command output, which is dumped to a text file, as a self refreshing regular top command.")
	fmt.Println("\nUsage: read-top [OPTIONS] FILENAME")
	fmt.Println("\nOptions")
	fmt.Println("\t-r,--refresh\t Normal Speed in seconds (Default 0.1)")
	fmt.Println("\t-s,--slow\t Slow Speed in seconds (Default 1). If entered smaller than 'Normal Speed', 'Normal Speed' is used")
	fmt.Println("\t-t,--token\t Token to Look (If the token is captured refresh rate will switch to 'slow')")
	fmt.Println("\t-l,--lines\t Number of lines to display (Default 50)")
	fmt.Println("\t-v,--version\t Application version")
	fmt.Println("\t-h,--help\t Display help")
	fmt.Println("\nExamples:")
	fmt.Println("\tread-top /tmp/top.out")
	fmt.Println("\tread-top -s 1 -l 25 /tmp/top.out")
	fmt.Println("\tread-top -s 0.2 -S 1 -t my_process /tmp/top.out")
}

func displayVersion() {
	fmt.Println(appVersion)
}
