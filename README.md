# sshcmd
`sshcmd` is a golang library that allows the user to execute a CLI command over the `ssh` protocol and retrieve its output.

## Example
```
logs, err := sshcmd.RemoteRun("ubuntu", "3.123.5.788", "/Users/jim/.ssh/id_rsa", "cat /home/ubuntu/server.log")
```
