package main

import "flag"

type AppFlags struct {
	CmdHelp    *bool
	InputFile  *string
	Driver     *string
	OutputFile *string
}

func initAppFlags() AppFlags {
	return AppFlags{
		CmdHelp:    flag.Bool("help", false, "Show usages"),
		InputFile:  flag.String("file", "", "Input file path of PlantUML syntax"),
		OutputFile: flag.String("out", "", "Output file path"),
		Driver:     flag.String("to", "", "Target database driver for generated query. Values are mariadb, mysql, postgresql"),
	}
}
