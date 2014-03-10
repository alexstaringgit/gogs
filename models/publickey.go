package models

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Unknwon/com"
)

var (
	//publicKeyRootPath string
	sshPath       string
	appPath       string
	tmplPublicKey = "### autogenerated by gitgos, DO NOT EDIT\n" +
		"command=\"%s serv key-%d\",no-port-forwarding," +
		"no-X11-forwarding,no-agent-forwarding,no-pty %s\n"
)

func exePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func homeDir() string {
	home, err := com.HomeDir()
	if err != nil {
		return "/"
	}
	return home
}

func init() {
	var err error
	appPath, err = exePath()
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}

	sshPath = filepath.Join(homeDir(), ".ssh")
}

type PublicKey struct {
	Id      int64
	OwnerId int64     `xorm:"index"`
	Name    string    `xorm:"unique not null"`
	Content string    `xorm:"text not null"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func GenAuthorizedKey(keyId int64, key string) string {
	return fmt.Sprintf(tmplPublicKey, appPath, keyId, key)
}

func AddPublicKey(key *PublicKey) error {
	_, err := orm.Insert(key)
	if err != nil {
		return err
	}

	err = SaveAuthorizedKeyFile(key)
	if err != nil {
		_, err2 := orm.Delete(key)
		if err2 != nil {
			// TODO: log the error
		}
		return err
	}

	return nil
}

func DeletePublicKey(key *PublicKey) error {
	_, err := orm.Delete(key)
	return err
}

func ListPublicKey(userId int64) ([]PublicKey, error) {
	keys := make([]PublicKey, 0)
	err := orm.Find(&keys, &PublicKey{OwnerId: userId})
	return keys, err
}

func SaveAuthorizedKeyFile(key *PublicKey) error {
	p := filepath.Join(sshPath, "authorized_keys")
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	//os.Chmod(p, 0600)
	_, err = f.WriteString(GenAuthorizedKey(key.Id, key.Content))
	return err
}
