package ioconn

import (
	"net"
	"time"
)

// ReaderWriterErorr represent error on both of reader and writer
type ReaderWriterErorr struct {
	Reader error
	Writer error
}

func (rwe ReaderWriterErorr) Error() string {
	rErr := ""
	if rwe.Reader != nil {
		rErr = rwe.Reader.Error()
	}
	wErr := ""
	if rwe.Writer != nil {
		wErr = rwe.Writer.Error()
	}
	return "reader: " + rErr + "; writer: " + wErr
}

type stringaddr struct {
	network string
	name    string
}

func (sa stringaddr) Network() string {
	return sa.network
}

func (sa stringaddr) String() string {
	return sa.name
}

type localaddr interface {
	LocalAddr() net.Addr
}

type remoteaddr interface {
	RemoteAddr() net.Addr
}

type setdeadline interface {
	SetDeadline(t time.Time) error
}

type setreaddeadline interface {
	SetReadDeadline(t time.Time) error
}

type setwritedeadline interface {
	SetWriteDeadline(t time.Time) error
}
