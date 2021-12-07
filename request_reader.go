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

func (r requestReader) ReadHeaders() (map[string][]string, error) {
	headers := make(map[string][]string)
	for {
		line, err := r.ReadLine()
		if err != nil {
			return headers, err
		}
		if string(line) == "" {
			return headers, nil
		}
		h := strings.Split(string(line), ":")
		headers[h[0]] = strings.Split(strings.TrimPrefix(h[1], " "), ",")
	}
}
