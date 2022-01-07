package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	userName       = "root"
	host           = "172.16.102.150"
	privateKeyFile = "/home/shea/.ssh/fuzamei/fuzamei_id_rsa"
	knownHostsFile = "/home/shea/.ssh/known_hosts"
)

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalln("error loading private key file", err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatalln("error parsing private key", err)
	}

	return privateKey
}

func checkServerPublicKey(host string, key ssh.PublicKey) error {
	knownHostFile, err := os.Open(knownHostsFile)
	if err != nil {
		log.Fatalln("error open known hosts file")
	}
	scanner := bufio.NewScanner(knownHostFile)
	var publicKey ssh.PublicKey

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}

		host = strings.Split(host, ":")[0]
		if strings.Contains(fields[0], host) {
			_, _, publicKey, _, _, err = ssh.ParseKnownHosts(scanner.Bytes())
			if err != nil {
				log.Fatalln("err parsing server public key", err)
			}

			if bytes.Equal(key.Marshal(), publicKey.Marshal()) {
				return nil
			}
		}
	}

	log.Println("key not exists in known_hosts")

	return nil
}

func hostKeyCallBack(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return checkServerPublicKey(hostname, key)
}

func main() {
	privateKey := getKeySigner(privateKeyFile)
	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: hostKeyCallBack,
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalln("error dialing server", err)
	}

	log.Println(string(client.ClientVersion()))

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("failed to create session:", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	err = session.RequestPty("vt100", 40, 80, ssh.TerminalModes{
		ssh.ECHO: 0,
	})

	if err != nil {
		log.Fatalln("error requesting psuedo-terminal", err)
	}

	err = session.Shell()
	if err != nil {
		log.Fatalln("err executing command", err)
	}

	_ = session.Wait()
}
