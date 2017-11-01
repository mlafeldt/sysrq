package main

import (
	"fmt"
	"os"

	"github.com/mlafeldt/sysrq"
)

func main() {
	var cmds []sysrq.Command
	for _, arg := range os.Args[1:] {
		cmd, err := sysrq.FromString(arg)
		if err != nil {
			abort("%s", err)
		}
		cmds = append(cmds, cmd)
	}

	if len(cmds) == 0 {
		abort("no command specified")
	}

	sys := sysrq.SysRq{TriggerFile: os.Getenv("TRIGGER_FILE")}

	for _, cmd := range cmds {
		fmt.Printf("Triggering SysRq command %s ...\n", cmd)
		if err := sys.Trigger(cmd); err != nil {
			abort("%s", err)
		}
	}
}

func abort(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", a...)
	os.Exit(1)
}
