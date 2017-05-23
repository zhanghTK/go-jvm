package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.isVersion {
		fmt.Println("version 0.0.1")
	} else if cmd.isHelp || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class: %s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
