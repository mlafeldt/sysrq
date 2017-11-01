# sysrq

[![GoDoc](https://godoc.org/github.com/mlafeldt/sysrq?status.svg)](https://godoc.org/github.com/mlafeldt/sysrq)

Go client to perform low-level commands via the [Linux SysRq interface](https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/sysrq.rst).

## Testing

Here's how to run SysRq commands against a local Vagrant box:

```bash
$ vagrant up

$ make trigger CMD=crash

$ make log
ubuntu-xenial login: [   94.116848] sysrq: SysRq : Trigger a crash
[   94.152571] BUG: unable to handle kernel NULL pointer dereference at           (null)
[   94.263679] IP: [<ffffffff81504df6>] sysrq_handle_crash+0x16/0x20
...
```
