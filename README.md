# sysrq

[![Build Status](https://travis-ci.org/mlafeldt/sysrq.svg?branch=master)](https://travis-ci.org/mlafeldt/sysrq)
[![GoDoc](https://godoc.org/github.com/mlafeldt/sysrq?status.svg)](https://godoc.org/github.com/mlafeldt/sysrq)

Go client to perform low-level commands via the [Linux SysRq interface](https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/sysrq.rst).

Among other things, SysRq can crash the system by forcing a NULL pointer dereference, which makes it a good fit for [Chaos Engineering experiments](https://medium.com/production-ready/chaos-engineering-101-1103059fae44).

[See GoDoc](https://godoc.org/github.com/mlafeldt/sysrq#Command) for a list of all supported commands.

## Playground

Here's how to run SysRq commands against a local Vagrant machine:

```bash
# Start Vagrant machine
$ vagrant up

# Trigger crash command
$ make trigger CMD=crash

# Show system logs
$ make log
ubuntu-xenial login: [   94.116848] sysrq: SysRq : Trigger a crash
[   94.152571] BUG: unable to handle kernel NULL pointer dereference at           (null)
[   94.263679] IP: [<ffffffff81504df6>] sysrq_handle_crash+0x16/0x20
...

# Fix Vagrant machine
$ vagrant reload
```
