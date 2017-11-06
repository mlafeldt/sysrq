# sysrq

[![](https://travis-ci.org/mlafeldt/sysrq.svg?branch=master)](https://travis-ci.org/mlafeldt/sysrq)
[![](https://img.shields.io/docker/pulls/mlafeldt/sysrq.svg?maxAge=604800)](https://hub.docker.com/r/mlafeldt/sysrq/)
[![](https://godoc.org/github.com/mlafeldt/sysrq?status.svg)](https://godoc.org/github.com/mlafeldt/sysrq)

Go client to perform low-level commands via the [Linux SysRq interface](https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/sysrq.rst) (accessible at `/proc/sysrq-trigger`).

Among other things, SysRq can crash the system by forcing a NULL pointer dereference, which makes it a good fit for [Chaos Engineering experiments](https://medium.com/production-ready/chaos-engineering-101-1103059fae44).

## CLI

In addition to the [Go library](https://godoc.org/github.com/mlafeldt/sysrq), there's a `sysrq` command-line tool you can install from source:

```bash
go get -u github.com/mlafeldt/sysrq/cmd/sysrq
```

Use the tool to trigger one or more commands:

```bash
sudo sysrq <cmd>...
```

This will print a list of all available commands:

```bash
sysrq -list
```

## Vagrant playground

Here's how to run SysRq commands against a local Vagrant machine:

```bash
# Start Vagrant machine
vagrant up

# Trigger crash command
make trigger CMD=crash

# Show system logs
make log
...
ubuntu-xenial login: [   94.116848] sysrq: SysRq : Trigger a crash
[   94.152571] BUG: unable to handle kernel NULL pointer dereference at           (null)
[   94.263679] IP: [<ffffffff81504df6>] sysrq_handle_crash+0x16/0x20
...

# Fix Vagrant machine
vagrant reload
```

## Docker

Since Docker mounts `/proc/sysrq-trigger` as read-only, you cannot run commands against other containers, but you can still affect the host system:

```bash
docker run --rm -v /proc/sysrq-trigger:/sysrq -e TRIGGER_FILE=/sysrq mlafeldt/sysrq <cmd>...
```

## Author

This project is being developed by [Mathias Lafeldt](https://twitter.com/mlafeldt).
