package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mlafeldt/sysrq"
)

func main() {
	listCommands := flag.Bool("list", false, "List SysRq commands")
	flag.Parse()

	if *listCommands {
		for _, cmd := range sysrq.Commands {
			fmt.Println(cmd)
		}
		return
	}

	if flag.NArg() == 0 {
		abort("no command specified")
	}

	var cmds []sysrq.Command
	for _, arg := range flag.Args() {
		cmd, err := sysrq.FromString(arg)
		if err != nil {
			abort("%s", err)
		}
		cmds = append(cmds, cmd)
	}

	sys := sysrq.SysRq{TriggerFile: os.Getenv("TRIGGER_FILE")}

	for _, cmd := range cmds {
		fmt.Printf("Triggering SysRq command %q ...\n", cmd)
		if err := sys.Trigger(cmd); err != nil {
			abort("%s", err)
		}
	}
}

func abort(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", a...)
	os.Exit(1)
}
