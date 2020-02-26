package ioconn

// ReaderWriterError represent error on both of reader and writer
type ReaderWriterError struct {
	Reader error
	Writer error
}

func (rwe ReaderWriterError) Error() string {
	rErr := ""
	if rwe.Reader != nil {
		rErr = rwe.Reader.Error()
	}
	wErr := ""
	if rwe.Writer != nil {
		wErr = rwe.Writer.Error()
	}
	ret := ""
	if rErr != "" {
		ret += "reader(" + rErr + ")"
	}
	if wErr != "" {
		if ret != "" {
			ret += ", "
		}
		ret += "writer(" + wErr + ")"
	}
	return ret
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
