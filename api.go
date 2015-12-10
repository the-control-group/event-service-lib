package lib

import (
	"bufio"
	"net"
	"net/textproto"
	"sync"
	"time"

	"gopkg.in/Sirupsen/logrus.v0"
)

// DEPRECATED because it's annoying
var ApiWelcomeMessage string = "This is the API welcome message"

var ApiMessageHandlerFns = []ApiMessageHandler{}

var ApiErrorHandlerFn ApiErrorHandler

type ApiMessageHandler func(msg []byte, writer *textproto.Writer)

type ApiErrorHandler func(error, *textproto.Writer)

func ListenWithAddress(log *logrus.Entry, addr Address, timeout time.Duration) (listener *net.TCPListener, err error) {
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
	go serve(log, listener, timeout)
	return
}

func Listen(log *logrus.Entry) (listener *net.TCPListener, err error) {
	address, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("", "10000"))
	if err != nil {
		log.WithError(err).Error("Unable to resolve tcp address")
		return
	}
	listener, err = net.ListenTCP("tcp", address)
	if err != nil {
		log.WithError(err).Error("Unable to establsih listener")
		return
	}
	go serve(log, listener, 60*time.Second)
	return
}

func serve(log *logrus.Entry, listener *net.TCPListener, timeout time.Duration) (err error) {
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
				handleConnection(log, c, timeout)
			}(c)
		}
	}
}

func handleConnection(log *logrus.Entry, c net.Conn, timeout time.Duration) {
	var err error
	defer c.Close()
	writer := textproto.NewWriter(bufio.NewWriter(c))
	// This is annoying
	// if ApiWelcomeMessage != "" {
	// 	writer.PrintfLine(ApiWelcomeMessage)
	// }
	bufc := bufio.NewReader(c)
	for {
		var msg []byte
		select {
		case <-Done:
			return
		default:
			c.SetDeadline(time.Now().Add(timeout))
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
			log.WithField("size", len(msg)).Debug("Read on connection")
			for _, h := range ApiMessageHandlerFns {
				h(msg, writer)
			}
		}
	}
}
