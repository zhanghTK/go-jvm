package main

import "flag"
import "fmt"
import "os"

// 命令行结构体
type Cmd struct {
	isHelp     bool
	isVersion  bool
	cpOption   string
	XjreOption string
	class      string
	args       []string
}

func parseCmd() *Cmd {
	// flag包帮助处理命令行选项
	cmd := &Cmd{}

	// 绑定函数，解析失败时用于提示
	flag.Usage = printUsage

	// 根据默认值绑定命令行选项到制定的变量
	flag.BoolVar(&cmd.isHelp, "help", false, "print help message")
	flag.BoolVar(&cmd.isHelp, "?", false, "print help message")
	flag.BoolVar(&cmd.isVersion, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	// 解析命令行参数
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	// os包的Args变量存放命令行的全部参数
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
