package Interfaces

import (
  "Requester/Objects"
  "Requester/Controllers"
)

type Request struct {
  Objects.Request
  Objects.Response 

}

func (r *Request) RequestURL() {

    res, err := Controllers.RequestURL(r.Method, r.URL)
    if err != nil {
      r.Err = err 
    }
    Response := &Objects.Response{
      Code: res.StatusCode,
      ContentLength: res.ContentLength, 
      RawBody: res.Body,
    }
    r.Response = *Response
}

func (r Request) DirectoryExist() {
  // Controllers.DirectoryExist(r.URL)
}

func (r Request) FileExist() {
  // Controllers.FileExist(r.URL)
}

func (r Request) CreateFile() {
  // Controllers.CreateFile()
}

func (r Request) ResWitter() {
  // Controllers.ResWitter(r)
}

