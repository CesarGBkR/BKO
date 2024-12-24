package Controllers

import (
  "net/http"
  "regexp"  
  "errors"
  // "fmt"
)

func RequestURL (url string)  (http.Response, error) {
    res, err := http.Get(url)
    return *res, err
}

func Extract(re regexp.Regexp, findOn string) ([][]string, error){
  match := re.FindAllStringSubmatch(findOn, -1)
  if len(match) == 0 {
    err := errors.New("No Matches Found")
    return match, err
  } 
  return match, nil
}  
