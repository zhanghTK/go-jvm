package main

import (
	"fmt"
	"GJvm/classpath"
	"strings"
)

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

// -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_91.jdk/Contents/Home/jre" java.lang.Object
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class: %s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Couldn't find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
