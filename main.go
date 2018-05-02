package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"gcg/gens/common"
	"log"
)

const (
	exec    = "gcg"
	version = "1.0.0"
)

var (
	command  string
	rootPath string
)

var (
	appName    string
	appPort    int
	ormFile    string
	tables	   string
)

var commandMap map[string]*Command

type CommandFuncType func(name, desc string)
type Command struct {
	Name string
	Desc string
	Func CommandFuncType
}

func init() {
	genutils.Asset = Asset
	// 初始化参数值
	flag.StringVar(&ormFile, "config", "config.yaml", "config file")
	flag.StringVar(&tables, "tables", "", "add table")
}

func showHelp(name, desc string) {
	commands := make([]string, 0, len(commandMap))
	for _, v := range commandMap {
		commands = append(commands, fmt.Sprintf("%s\t%s", v.Name, v.Desc))
	}
	printHelp(fmt.Sprintf("Usage: %s <command> [options]", exec), commands,
		[]string{
			"-config\tdatabase config file, default: config.yaml",
			"-tables\t db table name,split with ','",
		}, nil)
}

func initCommands() {
	for i, v := range os.Args {
		switch i {
		case 1:
			command = v
		}
	}

	// 初始化命令列表
	commandMap = map[string]*Command{
		"version": &Command{
			Name: "version",
			Desc: "show version",
			Func: showVersion,
		},
		"help": &Command{
			Name: "help",
			Desc: "show help",
			Func: showHelp,
		},
		"init": &Command{
			Name: "init",
			Desc: "init code from template \t[-config]",
			Func: genORM,
		},
		"add": &Command{
			Name: "add",
			Desc: "add code from template \t[-tables]",
			Func: genORM,
		},
	}
}

func checkArgs() bool {
	if command == "" {
		showHelp("help", commandMap["help"].Desc)
		return false
	} else if command == "gen_gin_server" {
		if appName == "" {
			log.Println("name can not be empty")
		}
		if appPort <= 0 {
			log.Println("port can not be empty")
		}
	}
	return true
}

func printHelp(usage string, commands, options, examples []string) {
	fmt.Println()
	fmt.Println(usage)
	if len(commands) > 0 {
		fmt.Println("\nCommands:\n")
		for _, s := range commands {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	if len(options) > 0 {
		fmt.Println("\nOptions:\n")
		for _, s := range options {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	if len(examples) > 0 {
		fmt.Println("\nExamples:\n")
		for _, s := range examples {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	fmt.Println()
}

func showVersion(name, desc string) {
	fmt.Println(version)
}

func parseRootDir() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s", err)
	}
	pathList := strings.Split(os.Getenv("GOPATH"), ":")
	for _, path := range pathList {
		if strings.HasPrefix(currentDir, path) {
			var prefix string
			if strings.HasSuffix(path, "/") {
				prefix = path + "src/"
			} else {
				prefix = path + "/src/"
			}
			currentDir = currentDir[len(prefix):]
		}

		rootPath = currentDir
	}
}
//go:generate go-bindata -o ./bindata.go -prefix "./template/" ./template/...
func main() {
	parseRootDir()
	fmt.Println(rootPath)
	initCommands()
	if len(os.Args) < 2 {
		showHelp("help", commandMap["help"].Desc)
		return
	}
	flag.CommandLine.Parse(os.Args[2:])
	if !checkArgs() {
		return
	}
	c := commandMap[command]
	if c == nil {
		showHelp("help", commandMap["help"].Desc)
		return
	} else {
		c.Func(c.Name, c.Desc)
	}
}
