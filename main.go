package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("plantsql: a tool that generates SQL queries from PlantUML syntax")

	// Parse flags
	appFlags := initAppFlags()
	flag.Parse()

	// Intercept command
	if *appFlags.CmdHelp {
		fmt.Println("Available Commands and Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Validate required options
	if v := appFlags.InputFile; v == nil || *v == "" {
		fmt.Println("-file is required")
		os.Exit(1)
	}
	if v := appFlags.Driver; v == nil || *v == "" {
		fmt.Println("-to is required")
		os.Exit(2)
	}

	fmt.Printf("DEBUG: InputFile = %s\n", *appFlags.InputFile)
	fmt.Printf("DEBUG: Driver = %s\n", *appFlags.Driver)

	// Init output file
	if v := appFlags.OutputFile; v == nil || *v == "" {
		fmt.Println("DEBUG: -out is not set. Create default output file name...")
		out := fmt.Sprintf("output-%s.sql", *appFlags.Driver)
		appFlags.OutputFile = &out
	}
	fmt.Printf("DEBUG: OutputFile = %s\n", *appFlags.OutputFile)
}
