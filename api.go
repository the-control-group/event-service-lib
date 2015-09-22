package lib

import (
	"bufio"
	"net"
	"net/textproto"
	"sync"
	"time"
)

var ApiWelcomeMessage string = "This is the API welcome message"

var ApiMessageHandlerFns = []ApiMessageHandler{}

var ApiErrorHandlerFn ApiErrorHandler

type ApiMessageHandler func(msg []byte, writer *textproto.Writer)

type ApiErrorHandler func(error, *textproto.Writer)

func Listen(log logger) (listener *net.TCPListener, err error) {
	// Automatically assign open port
	address, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if err != nil {
		log.Error("Unable to resolve tcp address", err)
		return
	}
	listener, err = net.ListenTCP("tcp", address)
	if err != nil {
		log.Error("Unable to establsih listener", err)
		return
	}
	go serve(log, listener)
	return
}

func serve(log logger, listener *net.TCPListener) (err error) {
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
				log.Debug(err)
				if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
					continue
				}
				return
			}
			wg.Add(1)
			go func(c *net.TCPConn) {
				defer wg.Done()
				handleConnection(log, c)
			}(c)
		}
	}
}

func handleConnection(log logger, c *net.TCPConn) {
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
				log.Debug("Reading bytes", err)
				if ApiErrorHandler != nil {
					ApiErrorHandler(err, writer)
				}
				continue
			}
			for _, h := range ApiMessageHandlers {
				h(msg, writer)
			}
		}
	}
}
