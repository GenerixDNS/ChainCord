package main

import (
	"fmt"
	"net"
	"strconv"
)

type channel struct {
	port     int
	host     string
	listener net.Listener
}

func (c channel) Open() ChannelImpl {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(c.port))
	c.listener = listener
	if err != nil {
		fmt.Println(err)
		return c
	}
	defer listener.Close()
	return c
}

func (c channel) Close() ChannelImpl {
	err := c.listener.Close()
	if err != nil {
		fmt.Println(c.listener.Close())
	}
	return c
}

type ChannelImpl interface {
	Open() ChannelImpl

	Close() ChannelImpl
}

func newChannel(port int, host string) ChannelImpl {
	c := channel{port: port, host: host}
	return c
}

func waitFor(stream net.Listener) (net.Conn, error) {
	con, err := stream.Accept()
	if err != nil {
		fmt.Println(err)
	}
	return con, err
}
