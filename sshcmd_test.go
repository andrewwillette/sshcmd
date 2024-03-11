package sshcmd

import (
	"fmt"
	"os"
	"testing"
)

var (
	PrivateKeyPath = os.Getenv("SSHCMD_TEST_PRIVATE_SSH_KEY")
	IPAddress      = os.Getenv("SSHCMD_TEST_SERVER_IP")
	User           = os.Getenv("SSHCMD_TEST_USER")
)

func TestRemoteRun(t *testing.T) {
	logs, err := RemoteRun(User, IPAddress, PrivateKeyPath, "ls -al")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("logs: %s\n", logs)
}
