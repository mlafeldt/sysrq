package main

import (
	"log"
	"os"

	"github.com/mlafeldt/sysrq"
)

func main() {
	var cmds []sysrq.Command
	for _, arg := range os.Args[1:] {
		cmd, err := sysrq.FromString(arg)
		if err != nil {
			log.Fatal(err)
		}
		cmds = append(cmds, cmd)
	}
	for _, cmd := range cmds {
		log.Printf("Triggering SysRq command %s ...\n", cmd)
		if err := sysrq.Trigger(cmd); err != nil {
			log.Fatal(err)
		}
	}
}
