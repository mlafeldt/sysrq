package sysrq_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mlafeldt/sysrq"
)

func TestFromString(t *testing.T) {
	want := sysrq.KillAllTasks
	got, err := sysrq.FromString("kill-all-tasks")
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}

	_, err1 := sysrq.FromString("foo")
	if err1 == nil || err1.Error() != "invalid command: foo" {
		t.Errorf("Expected error, got %v", err1)
	}
}

func TestString(t *testing.T) {
	want := "memory-full-oom-kill"
	got := sysrq.MemoryFullOOMKill.String()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestTrigger(t *testing.T) {
	want := sysrq.Sync

	sys := sysrq.SysRq{TriggerFile: "proc-sysrq-trigger"}
	if err := sys.Trigger(want); err != nil {
		t.Error(err)
	}
	defer os.Remove(sys.TriggerFile)

	got, err := ioutil.ReadFile(sys.TriggerFile)
	if err != nil {
		t.Error(err)
	}
	if len(got) != 1 || got[0] != byte(want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
