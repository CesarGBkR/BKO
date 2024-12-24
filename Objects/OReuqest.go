package Objects 

import (
  "net/http"
)

type Request struct {
  URL string
  Code int
  Response http.Response
  Err error
  Body string
  Dir string
}
