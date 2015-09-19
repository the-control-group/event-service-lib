package lib

import (
	"bufio"
	"io"
	"net"
	"net/textproto"
	"sync"
	"time"
)

func init() {
	for {
		var err = Listen()
		if Log != nil {
			Log.Error(err)
		}
		time.Sleep(5 * time.Second)
	}
}

var ApiMessageHandlers = []ApiMessageHandler{}

type ApiMessageHandler func(msg []byte, writer *textproto.Writer)

var ApiWelcomeMessage string = "This is the API welcome message"

var ApiErrorHandler = func(err error) {
	switch err.(type) {
	case net.Error:
		if err.(net.Error).Timeout() {
			if Log != nil {
				Log.Warn("Connection timed out", err)
			}
		}
	default:
		if err != io.EOF {
			if Log != nil {
				Log.Error("Reading bytes", err)
			}
		}
	}
	return
}

func Listen() (err error) {
	var listener *net.TCPListener
	// Automatically assign open port
	address, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if err != nil {
		if Log != nil {
			Log.Error("Reading bytes", err)
		}
		return
	}
	listener, err = net.ListenTCP("tcp", address)
	if err != nil {
		if Log != nil {
			Log.Error("Reading bytes", err)
		}
		return
	}
	defer listener.Close()
	serve(listener)
	return
}

func serve(listener *net.TCPListener) (err error) {
	var wg sync.WaitGroup
	for {
		select {
		case <-Done:
			wg.Wait()
			return
		default:
			var c *net.TCPConn
			c, err = listener.AcceptTCP()
			if err != nil {
				if Log != nil {
					Log.Error("Reading bytes", err)
				}
				if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
					continue
				}
				return
			}
			wg.Add(1)
			go func(c *net.TCPConn) {
				defer wg.Done()
				handleConnection(c)
			}(c)
		}
	}
}

func handleConnection(c *net.TCPConn) {
	var err error
	defer c.Close()
	var timeout = 60 * time.Second
	writer := textproto.NewWriter(bufio.NewWriter(c))
	writer.PrintfLine(ApiWelcomeMessage)
	for {
		var msg []byte
		select {
		case <-Done:
			return
		default:
			c.SetDeadline(time.Now().Add(timeout))
			bufc := bufio.NewReader(c)
			msg, err = bufc.ReadBytes('\n')
			if err != nil {
				if Log != nil {
					Log.Error("Reading bytes", err)
				}
				ApiErrorHandler(err)
				continue
			}
			for _, h := range ApiMessageHandlers {
				h(msg, writer)
			}
		}
	}
}
