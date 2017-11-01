// Package sysrq allows to perform low-level commands via the Linux SysRq interface.
// See https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/sysrq.rst
// and https://en.wikipedia.org/wiki/Magic_SysRq_key for more information.
package sysrq

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Command is a SysRq command.
type Command byte

const (
	// Reboot immediately reboots the system without syncing or unmounting
	// your disks.
	Reboot Command = 'b'

	// Crash performs a system crash by a NULL pointer dereference. A
	// crashdump will be taken if configured.
	Crash Command = 'c'

	// TerminateAllTasks sends a SIGTERM to all processes, except for init.
	TerminateAllTasks Command = 'e'

	// MemoryFullOOMKill calls the OOM killer to kill a memory hog process,
	// but doesn't panic if nothing can be killed.
	MemoryFullOOMKill Command = 'f'

	// KillAllTasks sends a SIGKILL to all processes, except for init.
	KillAllTasks Command = 'i'

	// ThawFilesystems forcibly "Just thaw it" - filesystems frozen by the
	// FIFREEZE ioctl.
	ThawFilesystems Command = 'j'

	// SAK (Secure Access Key) kills all programs on the current virtual
	// console.
	SAK Command = 'k'

	// ShowBacktraceAllActiveCPUs shows a stack backtrace for all active
	// CPUs.
	ShowBacktraceAllActiveCPUs Command = 'l'

	// ShowMemoryUsage dumps current memory info to your console.
	ShowMemoryUsage Command = 'm'

	// NiceAllRTTasks will make RT tasks nice-able.
	NiceAllRTTasks Command = 'n'

	// Poweroff shuts your system off (if configured and supported).
	Poweroff Command = 'o'

	// ShowRegisters dumps the current registers and flags to your console.
	ShowRegisters Command = 'p'

	// ShowAllTimers dumps per CPU lists of all armed hrtimers (but NOT
	// regular timer_list timers) and detailed information about all
	// clockevent devices.
	ShowAllTimers Command = 'q'

	// Unraw turns off keyboard raw mode and sets it to XLATE.
	Unraw Command = 'r'

	// Sync attempts to sync all mounted filesystems.
	Sync Command = 's'

	// ShowTaskStates dumps a list of current tasks and their information
	// to your console.
	ShowTaskStates Command = 't'

	// Unmount attempts to remount all mounted filesystems read-only.
	Unmount Command = 'u'

	// ShowBlockedTasks dumps tasks that are in uninterruptable (blocked)
	// state.
	ShowBlockedTasks Command = 'w'

	// DumpFtraceBuffer dumps the ftrace buffer.
	DumpFtraceBuffer Command = 'z'

	// Loglevel0 sets the console log level to 0. In this level, only
	// emergency messages like PANICs or OOPSes would make it to your
	// console.
	Loglevel0 Command = '0'

	// Loglevel1 sets the console log level to 1.
	Loglevel1 Command = '1'

	// Loglevel2 sets the console log level to 2.
	Loglevel2 Command = '2'

	// Loglevel3 sets the console log level to 3.
	Loglevel3 Command = '3'

	// Loglevel4 sets the console log level to 4.
	Loglevel4 Command = '4'

	// Loglevel5 sets the console log level to 5.
	Loglevel5 Command = '5'

	// Loglevel6 sets the console log level to 6.
	Loglevel6 Command = '6'

	// Loglevel7 sets the console log level to 7.
	Loglevel7 Command = '7'

	// Loglevel8 sets the console log level to 8.
	Loglevel8 Command = '8'

	// Loglevel9 sets the console log level to 9, the most verbose level.
	Loglevel9 Command = '9'

	// DefaultTriggerFile is the location of the file where SysRq commands
	// are written to by default.
	DefaultTriggerFile = "/proc/sysrq-trigger"
)

// FromString creates a Command from a string.
func FromString(s string) (Command, error) {
	cmd, ok := map[string]Command{
		"crash":                          Crash,
		"dump-ftrace-buffer":             DumpFtraceBuffer,
		"kill-all-tasks":                 KillAllTasks,
		"loglevel0":                      Loglevel0,
		"loglevel1":                      Loglevel1,
		"loglevel2":                      Loglevel2,
		"loglevel3":                      Loglevel3,
		"loglevel4":                      Loglevel4,
		"loglevel5":                      Loglevel5,
		"loglevel6":                      Loglevel6,
		"loglevel7":                      Loglevel7,
		"loglevel8":                      Loglevel8,
		"loglevel9":                      Loglevel9,
		"memory-full-oom-kill":           MemoryFullOOMKill,
		"nice-all-rt-tasks":              NiceAllRTTasks,
		"poweroff":                       Poweroff,
		"reboot":                         Reboot,
		"sak":                            SAK,
		"show-all-timers":                ShowAllTimers,
		"show-backtrace-all-active-cpus": ShowBacktraceAllActiveCPUs,
		"show-blocked-tasks":             ShowBlockedTasks,
		"show-memory-usage":              ShowMemoryUsage,
		"show-registers":                 ShowRegisters,
		"show-task-states":               ShowTaskStates,
		"sync":                           Sync,
		"terminate-all-tasks":            TerminateAllTasks,
		"thaw-filesystems":               ThawFilesystems,
		"unmount":                        Unmount,
		"unraw":                          Unraw,
	}[strings.ToLower(s)]
	if !ok {
		return 0, fmt.Errorf("invalid command: %s", s)
	}
	return cmd, nil
}

// String converts a Command to a string.
func (cmd Command) String() string {
	return map[Command]string{
		Crash:                      "crash",
		DumpFtraceBuffer:           "dump-ftrace-buffer",
		KillAllTasks:               "kill-all-tasks",
		Loglevel0:                  "loglevel0",
		Loglevel1:                  "loglevel1",
		Loglevel2:                  "loglevel2",
		Loglevel3:                  "loglevel3",
		Loglevel4:                  "loglevel4",
		Loglevel5:                  "loglevel5",
		Loglevel6:                  "loglevel6",
		Loglevel7:                  "loglevel7",
		Loglevel8:                  "loglevel8",
		Loglevel9:                  "loglevel9",
		MemoryFullOOMKill:          "memory-full-oom-kill",
		NiceAllRTTasks:             "nice-all-rt-tasks",
		Poweroff:                   "poweroff",
		Reboot:                     "reboot",
		SAK:                        "sak",
		ShowAllTimers:              "show-all-timers",
		ShowBacktraceAllActiveCPUs: "show-backtrace-all-active-cpus",
		ShowBlockedTasks:           "show-blocked-tasks",
		ShowMemoryUsage:            "show-memory-usage",
		ShowRegisters:              "show-registers",
		ShowTaskStates:             "show-task-states",
		Sync:                       "sync",
		TerminateAllTasks:          "terminate-all-tasks",
		ThawFilesystems:            "thaw-filesystems",
		Unmount:                    "unmount",
		Unraw:                      "unraw",
	}[cmd]
}

// SysRq is used to configure access to the Linux SysRq interface.
type SysRq struct {
	TriggerFile string
}

// Trigger performs one or more commands via the Linux SysRq interface.
func (sysrq SysRq) Trigger(cmd ...Command) error {
	if sysrq.TriggerFile == "" {
		sysrq.TriggerFile = DefaultTriggerFile
	}
	for _, c := range cmd {
		if err := ioutil.WriteFile(sysrq.TriggerFile, []byte{byte(c)}, 0644); err != nil {
			return err
		}
	}
	return nil
}

// Trigger performs one or more commands via the Linux SysRq interface.
func Trigger(cmd ...Command) error {
	return SysRq{}.Trigger(cmd...)
}
