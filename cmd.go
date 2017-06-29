package main

import (
	"flag"
	"fmt"
	"os"
)

// 接收控制台命令的结构体
type Cmd struct {
	isHelp     bool   // 接收-help
	isVersion  bool   // 接收-version
	cpOption   string // 接收-cp/-classpath
	XjreOption string
	class      string   // 接收主函数
	args       []string // 接收主函数参数
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage

	flag.BoolVar(&cmd.isHelp, "help", false, "print help message")
	flag.BoolVar(&cmd.isHelp, "?", false, "print help message")
	flag.BoolVar(&cmd.isVersion, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
