package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/nemith/netconf"
	ncssh "github.com/nemith/netconf/transport/ssh"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type netconfSessionArgs struct {
	host          string
	port          int
	timeout       int
	username      string
	password      string
	pubkey        string
	pubkeypass    string
	hostkeyignore bool
}

func BuildNetconfArgs(command *cobra.Command) netconfSessionArgs {
	host, _ := command.Flags().GetString("host")
	port, _ := command.Flags().GetInt("port")
	timeout, _ := command.Flags().GetInt("timeout")
	username, _ := command.Flags().GetString("username")
	password, _ := command.Flags().GetString("password")
	pubkey, _ := command.Flags().GetString("pubkey")
	pubkeypass, _ := command.Flags().GetString("pubkeypass")
	hostkeyignore, _ := command.Flags().GetBool("hostkeyignore")

	if username == "" {
		log.Fatal("username must be present for NETCONF session args")
	}
	if (password == "") && (pubkey == "") {
		log.Fatal("either password or pubkey must be present to setup NETCONF session args")
	}

	sessionArgs := netconfSessionArgs{
		host:          host,
		port:          port,
		timeout:       timeout,
		username:      username,
		password:      password,
		pubkey:        pubkey,
		pubkeypass:    pubkeypass,
		hostkeyignore: hostkeyignore,
	}

	return sessionArgs
}

func BuildSSHConfig(netconfArgs *netconfSessionArgs) ssh.ClientConfig {
	config := ssh.ClientConfig{User: netconfArgs.username}
	config.SetDefaults()
	config.Timeout = time.Second * 5

	if netconfArgs.pubkey != "" {
		// Read private key file (eg. ~/.ssh/id_rsa)
		pemBytes, err := os.ReadFile(netconfArgs.pubkey)
		if err != nil {
			log.Fatal(err)
		}

		var signer *ssh.Signer
		// Create signer to pass into auth method
		if netconfArgs.pubkeypass != "" {
			passPhrase := []byte(netconfArgs.pubkeypass)
			*signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, passPhrase)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			*signer, err = ssh.ParsePrivateKey(pemBytes)
			if err != nil {
				log.Fatal(err)
			}
		}

		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(*signer)}

	} else if netconfArgs.password != "" {
		config.Auth = []ssh.AuthMethod{ssh.Password(netconfArgs.password)}
	} else {
		log.Fatal("unable to setup ssh client config due to missing pubkey or password")
	}

	if netconfArgs.hostkeyignore {
		config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	return config
}

func SetupNetconfSession(sessionArgs netconfSessionArgs, sshConfig *ssh.ClientConfig) *netconf.Session {

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	transport, err := ncssh.Dial(ctx, "tcp", fmt.Sprintf("%v:%v", sessionArgs.host, sessionArgs.port), sshConfig)
	if err != nil {
		log.Fatal(err)
	}

	session, err := netconf.Open(transport)
	if err != nil {
		log.Fatal(err)
	}

	return session
}

type Capability struct {
	scheme    string
	authority string
	path      string
	query     string
	fragment  string
}

func BuildCapabilitiesMap(uris []string) (*[]url.URL, error) {
	capabilities := make([]url.URL, len(uris))

	for i, u := range uris {
		urlStr, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		capabilities[i] = *urlStr
	}

	return &capabilities, nil
}
