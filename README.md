# sshcmd
`sshcmd` is a golang library that allows the user to execute a CLI command over the `ssh` protocol and retrieve its output.

## Example
```
logs, err := sshcmd.RemoteRun(User, IPAddress, PrivateKeyPath, "cat "+LogLocation)
```
