package sshcmd

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// RemoteRun runs a command on a remote server and returns its output
func RemoteRun(user, ip, privateKeyFilePath, command string) (string, error) {
	privateKey, err := getFileAsString(privateKeyFilePath)
	if err != nil {
		return "", err
	}
	// run command on remote server
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort(ip, "22"), config)
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(command)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func getFileAsString(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("sshcmd.getFileAsString error: %s", err)
	}
	return string(file), nil
}
