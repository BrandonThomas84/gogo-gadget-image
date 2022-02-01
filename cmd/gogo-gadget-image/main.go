package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BrandonThomas84/gogo-gadget-image/internal"
	"github.com/jessevdk/go-flags"
)

func main() {

	commands, err := internal.GetImageProcessorCommands()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	fmt.Printf("\n\nImage Processor Configuration: \n")
	fmt.Printf("%s\n\n", strings.Repeat("=", 50))
	fmt.Printf("Input Directory: %v\n", commands.Directory)
	fmt.Printf("Output Directory: %v\n", commands.OutputDirectory)

	fmt.Println("\n\nprocessing....")

	// here is where we would want to take a look at the the different
	// options that were passed in and use them to make decisions on
	// how to process the images
}
