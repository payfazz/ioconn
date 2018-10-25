package ioconn

import (
	"io"
)

// Config is parameter to New
type Config struct {
	Writer     io.Writer
	Reader     io.Reader
	LocalName  string
	RemoteName string
}
