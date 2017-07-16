package main

func main() {
	cmd := parseCmd()

	if cmd.isVersion {
		println("version 0.0.1")
	} else if cmd.isHelp || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
