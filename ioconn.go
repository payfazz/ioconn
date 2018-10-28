package ioconn

import (
	"fmt"
	"io"
	"net"
	"time"
)

type ioconn struct {
	io.Reader
	io.Writer
	localaddr  net.Addr
	remoteaddr net.Addr
}

// New net.Conn based on Config
func New(config Config) net.Conn {
	ret := &ioconn{
		Reader: config.Reader,
		Writer: config.Writer,
	}
	if item, ok := ret.Writer.(localaddr); ok && config.LocalName == "" {
		ret.localaddr = item.LocalAddr()
	} else {
		ret.localaddr = stringaddr{"io.Writer", config.LocalName}
	}
	if item, ok := ret.Reader.(remoteaddr); ok && config.LocalName == "" {
		ret.remoteaddr = item.RemoteAddr()
	} else {
		ret.remoteaddr = stringaddr{"io.Reader", config.RemoteName}
	}
	return ret
}

func (ic *ioconn) Close() error {
	ret := ReaderWriterError{}
	if item, ok := ic.Reader.(io.Closer); ok {
		ret.Reader = item.Close()
	}
	if item, ok := ic.Writer.(io.Closer); ok {
		ret.Writer = item.Close()
	}
	if ret.Reader == nil && ret.Writer == nil {
		return nil
	}
	return ret
}

func (ic *ioconn) LocalAddr() net.Addr {
	return ic.localaddr
}

func (ic *ioconn) RemoteAddr() net.Addr {
	return ic.remoteaddr
}

func (ic *ioconn) SetDeadline(t time.Time) error {
	ret := ReaderWriterError{}
	if item, ok := ic.Reader.(setdeadline); ok {
		ret.Reader = item.SetDeadline(t)
	} else {
		ret.Reader = fmt.Errorf("reader doesn't implement SetDeadLine")
	}
	if item, ok := ic.Writer.(setdeadline); ok {
		ret.Writer = item.SetDeadline(t)
	} else {
		ret.Writer = fmt.Errorf("writer doesn't implement SetDeadLine")
	}
	if ret.Reader == nil && ret.Writer == nil {
		return nil
	}
	return ret
}

func (ic *ioconn) SetReadDeadline(t time.Time) error {
	if item, ok := ic.Reader.(setreaddeadline); ok {
		return item.SetReadDeadline(t)
	}
	return fmt.Errorf("reader doesn't implement SetReadDeadline")
}

func (ic *ioconn) SetWriteDeadline(t time.Time) error {
	if item, ok := ic.Reader.(setwritedeadline); ok {
		return item.SetWriteDeadline(t)
	}
	return fmt.Errorf("reader doesn't implement SetWriteDeadline")
}
