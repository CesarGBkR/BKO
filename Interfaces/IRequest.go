package Interfaces

import (
  "Requester/Objects"
  "Requester/Controllers"
)

type Request struct {
  Objects.Request

}

func (r Request) RequestURL() {
    res, err := Controllers.RequestURL(r.URL)
    if err != nil {
      r.Err = err
    }
    r.Code = res.StatusCode
    r.Response = res

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
