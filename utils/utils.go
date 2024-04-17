package utils

import (
	"bytes"
	"cobra-Azure-tools/database"
	"fmt"
	"net"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"

	"log"
)

func Show(name string, age int) {
	fmt.Printf("name is %s, age is %d", name, age)
}

func sshConnect(user, pwd, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(pwd))

	// host key callbk
	hostKeyCallbk := func(host string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: hostKeyCallbk,
		BannerCallback:  nil,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func Sysinfo(host, pwd, username string, port int, cmd string, excel bool) {
	db := database.ContentMysql()
	var stdOut, stdErr bytes.Buffer
	// 使用用户名，密码登陆
	session, err := sshConnect(username, pwd, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run(cmd)
	message := strings.Replace(stdOut.String(), "\n", " ", -1)
	fmt.Println(message)

	data := strings.Split(message, " ")

	if cmd == "uptime" {
		node := database.NodeInfo{}
		node.IP = host

		node.Datetime = data[1]
		node.Status = data[2]
		node.StartedTime = strings.ReplaceAll(data[4], ",", "")
		node.UserNumber, _ = strconv.Atoi(data[6])
		node.LoadAverage1m, _ = strconv.ParseFloat(strings.ReplaceAll(data[11], ",", ""), 64)
		node.LoadAverage5m, _ = strconv.ParseFloat(strings.ReplaceAll(data[12], ",", ""), 64)
		node.LoadAverage15m, _ = strconv.ParseFloat(data[13], 64)

		_ = db.Model(&database.NodeInfo{}).Save(&node)
	}

}
