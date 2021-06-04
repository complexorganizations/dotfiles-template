package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

var (
	gitConfigContent = ``
	sshConfigContent = ``
	gpgKeyContent    = ``
	sshKeyContent    = ``
)

var (
	gitConfig     = fmt.Sprint(userDirectory() + "/.gitconfig")
	keysDir       = fmt.Sprint(userDirectory() + "/.ssh/")
	sshConfig     = fmt.Sprint(keysDir + "config")
	privateSSHKey = fmt.Sprint(keysDir + "id_ssh")
	privateGPGKey = fmt.Sprint(keysDir + "id_gpg")
	err           error
)

func init() {
	if !commandExists("git") {
		log.Fatal("Error: Please install git")
	}
	if !commandExists("gpg") {
		log.Fatal("Error: Please install gpg")
	}
}

func main() {
	installSSHKeys()
}

func installSSHKeys() {
	if !folderExists(keysDir) {
		err = os.Mkdir(keysDir, 0700)
		handleErrors(err)
	}
	if len(gitConfigContent) > 1 && len(gitConfig) > 1 {
		err = os.WriteFile(gitConfig, []byte(gitConfigContent), 0600)
		handleErrors(err)
	}
	if len(sshConfigContent) > 1 && len(sshConfig) > 1 {
		err = os.WriteFile(sshConfig, []byte(sshConfigContent), 0600)
		handleErrors(err)
	}
	if len(gpgKeyContent) > 1 && len(privateGPGKey) > 1 {
		err = os.WriteFile(privateGPGKey, []byte(gpgKeyContent), 0600)
		handleErrors(err)
	}
	if len(sshKeyContent) > 1 && len(privateSSHKey) > 1 {
		err = os.WriteFile(privateSSHKey, []byte(sshKeyContent), 0600)
		handleErrors(err)
	}
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func userDirectory() string {
	user, err := user.Current()
	handleErrors(err)
	return user.HomeDir
}

func commandExists(cmd string) bool {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = cmd
	return true
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
