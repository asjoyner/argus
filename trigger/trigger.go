package trigger

import (
	"bytes"
	"log"
	"net"
)

var (
	port     = 27487
	sentinel = []byte("Zeus is here!")
)

type Watcher struct {
	conn *net.UDPConn
}

func NewWatcher() (*Watcher, error) {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	})
	if err != nil {
		return nil, err
	}
	return &Watcher{conn: socket}, nil
}

func (w *Watcher) Wait() {
	for {
		data := make([]byte, len(sentinel))
		_, remoteAddr, err := w.conn.ReadFromUDP(data)
		if err != nil {
			log.Println("error reading from socket")
			continue
		}
		if bytes.Equal(data, sentinel) {
			return
		}
		log.Println("received malformed packet from %+v: %+v", remoteAddr, data)
	}
}

func Trigger() error {
	dest := &net.UDPAddr{IP: net.IPv4bcast, Port: port}
	socket, err := net.DialUDP("udp4", nil, dest)
	if err != nil {
		return err
	}

	if _, err := socket.Write(sentinel); err != nil {
		return err
	}
	return nil
}
