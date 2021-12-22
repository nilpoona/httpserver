package main

import "io"

type (
	Request struct {
		Method string
		Header map[string][]string
		Body   io.ReadCloser
	}
)

const (
	HeaderNameContentLength = "Content-Length"
)
