package Objects 

import (
  "io"
)

type Request struct {

  Method string 
  URL string
	Host string
  Code int
  Response Response 
	// Headers  map[string]string
  Body string
  Dir string
  Err error
}

type Response struct {
	Code    int
	Headers       map[string][]string
	ContentLength int64
	ContentWords  int64
	ContentLines  int64
	ContentType   string
  RawBody io.ReadCloser
}
