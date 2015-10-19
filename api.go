package lib

import (
	"bufio"
	"net"
	"net/textproto"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
)

var ApiWelcomeMessage string = "This is the API welcome message"

var ApiMessageHandlerFns = []ApiMessageHandler{}

var ApiErrorHandlerFn ApiErrorHandler

type ApiMessageHandler func(msg []byte, writer *textproto.Writer)

type ApiErrorHandler func(error, *textproto.Writer)

func ListenWithAddress(log *logrus.Entry, addr Address) (listener *net.TCPListener, err error) {
	// Automatically assign open port
	address, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(addr.Host, addr.Port))
	if err != nil {
		log.WithError(err).Error("Unable to resolve tcp address")
		return
	}
	listener, err = net.ListenTCP("tcp", address)
	if err != nil {
		log.WithError(err).Error("Unable to establsih listener")
		return
	}
	go serve(log, listener)
	return
}

func Listen(log *logrus.Entry) (listener *net.TCPListener, err error) {
	// Automatically assign open port
	address, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if err != nil {
		log.WithError(err).Error("Unable to resolve tcp address")
		return
	}
	listener, err = net.ListenTCP("tcp", address)
	if err != nil {
		log.WithError(err).Error("Unable to establsih listener")
		return
	}
	go serve(log, listener)
	return
}

func serve(log *logrus.Entry, listener *net.TCPListener) (err error) {
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
				log.WithError(err).Debug("Unable to accept TCP")
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

func handleConnection(log *logrus.Entry, c *net.TCPConn) {
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
				if err.Error() != "EOF" {
					log.WithError(err).Debug("Unable to read from connection")
					if ApiErrorHandlerFn != nil {
						ApiErrorHandlerFn(err, writer)
					}
				}
				return
			}
			for _, h := range ApiMessageHandlerFns {
				h(msg, writer)
			}
		}
	}
}
