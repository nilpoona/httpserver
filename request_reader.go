package main

import (
	"bufio"
	"io"
	"strings"
)

type (
	requestReader struct {
		bufr *bufio.Reader
	}
)

func newRequestReader(r io.Reader) *requestReader {
	return &requestReader{bufr: bufio.NewReader(r)}
}

func parseRequestLine(requestLine string) (method, requestURI, httpVersion string) {
	// https://datatracker.ietf.org/doc/html/rfc2616#section-5.1
	s := strings.Split(requestLine, " ")
	if len(s) < 3 {
		return
	}
	return s[0], s[1], s[2]
}

func (r *requestReader) ReadLine() ([]byte, error) {
	var line []byte
	for {
		l, isPrefix, err := r.bufr.ReadLine()
		if err != nil {
			return nil, err
		}
		if l != nil && !isPrefix {
			return l, nil
		}
		line = append(line, l...)
		if !isPrefix {
			break
		}
	}
	return line, nil
}
