package ssh

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

type SshConfig struct {
	Key      string
	Host     string
	Port     string
	User     string
	Password string
}

func OpenSSH(conf *SshConfig) (*ssh.Client, error) {
	// error handring
	eh := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}
	auth := []ssh.AuthMethod{}

	if conf.Key != "" {
		key, err := ioutil.ReadFile(conf.Key)
		eh(err, "private key")
		signer, err := ssh.ParsePrivateKey(key)
		eh(err, "signer")

		auth = append(auth, ssh.PublicKeys(signer))
	}

	if conf.Password != "" {
		auth = append(auth, ssh.Password(conf.Password))
	}

	config := &ssh.ClientConfig{
		User:            conf.User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", conf.Host+":"+conf.Port, config)
}
